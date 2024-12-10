package main

import (
	"fmt"
	"time"

	"ponchik/vectors"
)

func main() {
	width := 191
	height := 60
	aspect := float32(width) / float32(height)
	pixelAspect := float32(11) / 24
	gradient := []rune(" .:!/r(l1Z4H9W8$@")
	gradientSize := float32(len(gradient) - 2)
	screen := make([]rune, width*height)

	for t := 0; t < 10000; t++ {
		light := vectors.Norm(vectors.InitValuesVec3(-0.5, 0.5, -1.0))
		spherePos := vectors.InitValuesVec3(0, 3, 0)

		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {
				uv := (vectors.InitValuesVec2(float32(i), float32(j)).
					Div(vectors.InitValuesVec2(float32(width), float32(height))).
					Mult(vectors.InitValueVec2(2.0))).
					Minus(vectors.InitValueVec2(1.0))
				uv.X *= aspect * pixelAspect

				rayOrigin := vectors.InitValuesVec3(-6, 0, 0)
				rayDir := vectors.Norm(vectors.InitValuesVec3(2.0, uv.X, uv.Y))

				rayOrigin = vectors.RotateY(rayOrigin, 0.25)
				rayDir = vectors.RotateY(rayDir, 0.25)
				rayOrigin = vectors.RotateZ(rayOrigin, float64(t)*0.01)
				rayDir = vectors.RotateZ(rayDir, float64(t)*0.01)

				diff := float32(1.0)

				for k := 0; k < 5; k++ {
					minIt := float32(99999.0)
					intersection := vectors.Sphere(rayOrigin.Minus(spherePos), rayDir, 1)
					n := vectors.InitValueVec3(0)
					albedo := float32(1)

					// responsible for circle
					if intersection.X > 0 {
						itPoint := rayOrigin.Minus(spherePos).Plus(rayDir.Mult(vectors.InitValueVec3(intersection.X)))
						minIt = intersection.X
						n = vectors.Norm(itPoint)
					}
					// responsible for box
					boxN := vectors.InitValueVec3(0)
					intersection = vectors.Box(rayOrigin, rayDir, vectors.InitValueVec3(1), &boxN)
					if intersection.X > 0 && intersection.X < minIt {
						minIt = intersection.X
						n = boxN
					}

					// responsible for shadow
					intersection = vectors.InitValuesVec2(
						vectors.Plane(rayOrigin, rayDir, vectors.InitValuesVec3(0.0, 0.0, -1.0), 1.0),
						1.0)
					if intersection.X > 0 && intersection.X < minIt {
						minIt = intersection.X
						n = vectors.InitValuesVec3(0, 0, -1)
						albedo = 0.5
					}

					if minIt < 99999 {
						diff *= albedo * (float32(vectors.Dot(n, light)*0.5 + 0.5))
						rayOrigin = rayOrigin.Plus(rayDir.Mult(vectors.InitValueVec3(minIt - 0.01)))
						rayDir = vectors.Reflect(rayDir, n)
					} else {
						break
					}
				}

				color := diff * 20
				color = vectors.Clamp(color, 0, gradientSize)
				pixel := gradient[int(color)]
				// if uv.X*uv.X+uv.Y*uv.Y > 0.1 {
				// pixel = ' '
				// }
				screen[i+j*width] = pixel
			}
		}

		// Convert screen array to string
		screenOutput := ""
		for j := 0; j < height; j++ {
			screenOutput += string(screen[j*width:(j+1)*width]) + "\n"
		}

		// // Clear the terminal
		// fmt.Print("\033[H\033[2J")

		// Print the screen output
		fmt.Print(screenOutput)

		// Add a small delay for animation
		time.Sleep(50 * time.Millisecond)
	}
}
