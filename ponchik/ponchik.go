package main

import (
	"ponchik/vectors"
)

func main() {
	width := 120 * 2
	height := 30 * 2

	aspect := float32(width / height)
	pixelAspect := float32(11 / 24)
	gradient := []rune(" .:!/r(l1Z4H9W8$@")
	gradientSize := len(gradient) - 2

	screen := make([]rune, width*height)

	for t := 0; t < 10000; t++ {
		light := vectors.Norm(vectors.initValuesVec3(-0.5, 0.5, -1.0))
		spherePos := vectors.initValuesVec3(0, 3, 0)
		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {
				uv := vectors.initValuesVec2(i, j).Div(vectors.initValuesVec2(width, height)).Mult(2).Minus(1)
				uv.x *= aspect * pixelAspect
				rayOrigin := vectors.initValuesVec3(-6, 0, 0)
				rayDir := vectors.Norm(vectors.initValueVec(2, uv))
				rayOrigin = vectors.RotateY(rayOrigin, 0.25)
				rayDir = vectors.RotateY(rayDir, 0.25)
				rayOrigin = vectors.RotateZ(rayOrigin, float32(t)*0.01)
				rayDir = vectors.RotateZ(rayDir, float32(t)*0.01)
				diff := 1.0
				for k := 0; k < 5; k++ {
					minIt := 99999.0
					intersection := vectors.Sphere(rayOrigin.Minus(spherePos), rayDir, 1)
					n := vectors.initValueVec3(0)
					albedo := 1.0
					if intersection.x > 0 {
						itPoint := rayOrigin.Minus(spherePos).Plus(rayDir.Mult(vectors.initValueVec3(intersection.x)))
						minIt = intersection.x
						n = vectors.Norm(itPoint)
					}
					boxN := vectors.initValueVec3(0)
					intersection = vectors.Box(rayOrigin, rayDir, vectors.initValueVec3(1), boxN)
					if intersection.x > 0 && intersection.x < minIt {
						minIt = intersection.x
						n = boxN
					}

					intersection = vectors.Box(rayOrigin, rayDir, vectors.initValuesVec3(0, 0, -1), 1)
					if intersection.x > 0 && intersection.x < minIt {
						minIt = intersection.x
						n = vectors.initValuesVec3(0, 0, -1)
						albedo = 0.5
					}
					if minIt < 99999 {
						diff *= albedo * (vectors.Dot(n, light)*0.5 + 0.5)
						rayOrigin = rayOrigin.Plus(rayDir.Mult(vectors.initValueVec3(minIt - 0.01)))
						rayDir = vectors.Reflect(rayDir, n)
					} else {
						break
					}
				}
				color := int(diff * 20)
				color = vectors.Clamp(color, 0, gradientSize)
				pixel := gradient[color]
				screen[i+j*width] = pixel
			}
		}
		screen[width*height-1] = 0x00

	}
}
