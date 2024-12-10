package main

import (
	"fmt"
	"math"
	"time"

	"ponchik/vectors"
)

// Function to calculate the signed distance from a point to a torus
func GetDist(p *vectors.Vec3, t float64) float64 {
	// Rotate the point `p` around the Z-axis using a 2D rotation matrix
	rotatedX := p.X*math.Cos(t) - p.Z*math.Sin(t)
	rotatedZ := p.X*math.Sin(t) + p.Z*math.Cos(t)
	rotatedP := vectors.InitValuesVec3(rotatedX, p.Y, rotatedZ)

	// Rotate the point `rotatedP` further around the X-axis
	rotatedY := rotatedP.Y*math.Cos(t) - rotatedP.Z*math.Sin(t)
	rotatedZ = rotatedP.Y*math.Sin(t) + rotatedP.Z*math.Cos(t)
	rotatedP = vectors.InitValuesVec3(rotatedP.X, rotatedY, rotatedZ)

	// Calculate the distance from the rotated point to the toroidal shape
	// Toroidal shape is defined with a major radius of 1.0 and minor radius of 0.5
	q := vectors.InitValuesVec2(
		vectors.Len2(vectors.InitValuesVec2(rotatedP.X, rotatedP.Y))-1.0, // Offset for the major radius
		rotatedP.Z, // Minor radius along the Z-axis
	)
	return vectors.Len2(q) - 0.5 // Adjust for the minor radius
}

func main() {
	// Screen dimensions
	width := 191.0             // Number of columns
	height := 60.0             // Number of rows
	aspect := width / height   // Aspect ratio of the screen
	pixelAspect := 11.0 / 24.0 // Pixel aspect ratio to correct distortion

	// Gradient characters used to represent brightness levels
	gradient := []rune(" .:!/r(l1Z4H9W8$@")
	gradientSize := float64(len(gradient) - 2)

	// Screen buffer to store the characters
	screen := make([]rune, int(width*height))

	// Animation loop
	for t := 0.0; ; t += 0.05 {
		// Direction of the light source (normalized vector)
		light := vectors.Norm(vectors.InitValuesVec3(-0.5, 1.0, -0.5))

		// Iterate over each pixel in the screen
		for i := 0.0; i < width; i++ {
			for j := 0.0; j < height; j++ {
				// Map screen coordinates to normalized device coordinates (NDC)
				uv := vectors.InitValuesVec2(i, j).
					Div(vectors.InitValuesVec2(width, height)). // Scale to 0-1 range
					Mult(vectors.InitValueVec2(2.0)).           // Scale to -1 to 1 range
					Minus(vectors.InitValueVec2(1.0))           // Center to origin

				uv.X *= aspect * pixelAspect // Adjust for aspect and pixel ratios

				// Define the ray starting position and direction
				rayOrigin := vectors.InitValuesVec3(-2, 0, 0)                   // Camera position
				rayDir := vectors.Norm(vectors.InitValuesVec3(1.0, uv.X, uv.Y)) // Normalized ray direction

				// Ray marching parameters
				minDist := 0.01 // Minimum distance to consider a hit
				maxDist := 10.0 // Maximum distance for the ray
				dist := 0.0     // Accumulated distance along the ray
				iters := 0      // Iteration counter
				maxIters := 100 // Maximum number of ray marching steps
				hit := false    // Flag to check if a surface was hit

				// Ray marching loop
				for iters < maxIters {
					point := rayOrigin.Plus(rayDir.Mult(vectors.InitValueVec3(dist))) // Current point on the ray
					d := GetDist(point, t)                                            // Distance to the torus
					if d < minDist {                                                  // Hit detected
						hit = true
						break
					}
					if dist > maxDist { // Exceeded maximum distance
						break
					}
					dist += d // Advance the ray
					iters++
				}

				// Determine the pixel color based on the hit
				color := 0
				pixel := ' '
				if hit {
					// Recalculate the point and distance
					point := rayOrigin.Plus(rayDir.Mult(vectors.InitValueVec3(dist)))
					d := GetDist(point, t)

					// Calculate the surface normal at the hit point
					normal := vectors.Norm(vectors.InitValuesVec3(
						GetDist(&vectors.Vec3{X: point.X + 0.01, Y: point.Y, Z: point.Z}, t)-d,
						GetDist(&vectors.Vec3{X: point.X, Y: point.Y + 0.01, Z: point.Z}, t)-d,
						GetDist(&vectors.Vec3{X: point.X, Y: point.Y, Z: point.Z + 0.01}, t)-d,
					))

					// Compute diffuse lighting based on the angle between normal and light
					diff := math.Max(vectors.Dot(normal, light), 0.1) // Minimum brightness of 0.1
					color = int(diff * 25)                            // Map brightness to gradient range
				}

				// Clamp the color index and map to a character in the gradient
				color = int(vectors.Clamp(float64(color), 0, gradientSize))
				pixel = gradient[color]

				// Write the pixel to the screen buffer
				screen[int(i+j*width)] = pixel
			}
		}

		// Convert screen buffer to a string
		screenOutput := ""
		for j := 0; j < int(height); j++ {
			screenOutput += string(screen[j*int(width):(j+1)*int(width)]) + "\n"
		}

		// Clear the terminal for smooth animation
		fmt.Print("\033[H\033[2J")

		// Display the frame
		fmt.Print(screenOutput)

		// Add a delay to control the frame rate (approximately 30 FPS)
		time.Sleep(33 * time.Millisecond)
	}
}
