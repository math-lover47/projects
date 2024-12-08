package vectors

import (
	"fmt"
	"math"
)

type vec3 struct {
	x, y, z float32
}

// Initialize all components to the same value
func initValueVec3(val float32) *vec3 {
	return &vec3{
		x: val,
		y: val,
		z: val,
	}
}

func initValueVec(val float32, v2 *vec2) *vec3 {
	return &vec3{
		x: val,
		y: v2.x,
		z: v2.y,
	}
}

func initValuesVec3(x_, y_, z_ float32) *vec3 {
	return &vec3{
		x: x_,
		y: y_,
		z: z_,
	}
}

func (v *vec3) OppositeSign() *vec3 {
	return &vec3{x: -v.x, y: -v.y, z: -v.z}
}

// Add two vectors (returns a new vector)
func (v *vec3) Plus(other *vec3) *vec3 {
	return &vec3{
		x: v.x + other.x,
		y: v.y + other.y,
		z: v.z + other.z,
	}
}

// Subtract two vectors (returns a new vector)
func (v *vec3) Minus(other *vec3) *vec3 {
	return &vec3{
		x: v.x - other.x,
		y: v.y - other.y,
		z: v.z - other.z,
	}
}

// Multiply two vectors (returns a new vector)
func (v *vec3) Mult(other *vec3) *vec3 {
	return &vec3{
		x: v.x * other.x,
		y: v.y * other.y,
		z: v.z * other.z,
	}
}

// Divide two vectors (returns a new vector)
func (v *vec3) Div(other *vec3) *vec3 {
	// Check for division by zero
	if other.x == 0 || other.y == 0 || other.z == 0 {
		fmt.Println("Error: Division by zero in vector components")
		return &vec3{} // Return a zero vector as fallback
	}
	return &vec3{
		x: v.x / other.x,
		y: v.y / other.y,
		z: v.z / other.z,
	}
}

func Len3(v *vec3) float32 {
	return float32(math.Sqrt(float64(v.x*v.x + v.y*v.y + v.z*v.z)))
}
