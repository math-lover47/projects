package vectors

import (
	"fmt"
	"math"
)

type Vec3 struct {
	X, Y, Z float64
}

// InitValueVec3 initializes all components to the same value
func InitValueVec3(val float64) *Vec3 {
	return &Vec3{
		X: val,
		Y: val,
		Z: val,
	}
}

// InitValuesVec3 initializes vector with specific X, Y, and Z values
func InitValuesVec3(X, Y, Z float64) *Vec3 {
	return &Vec3{
		X: X,
		Y: Y,
		Z: Z,
	}
}

// OppositeSign returns a vector with opposite signs
func (v *Vec3) OppositeSign() *Vec3 {
	return &Vec3{X: -v.X, Y: -v.Y, Z: -v.Z}
}

// Plus adds two vectors (returns a new vector)
func (v *Vec3) Plus(other *Vec3) *Vec3 {
	return &Vec3{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

// Minus subtracts two vectors (returns a new vector)
func (v *Vec3) Minus(other *Vec3) *Vec3 {
	return &Vec3{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}

// Mult multiplies two vectors (returns a new vector)
func (v *Vec3) Mult(other *Vec3) *Vec3 {
	return &Vec3{
		X: v.X * other.X,
		Y: v.Y * other.Y,
		Z: v.Z * other.Z,
	}
}

// Div divides two vectors (returns a new vector)
func (v *Vec3) Div(other *Vec3) *Vec3 {
	// Check for division by zero
	if other.X == 0 || other.Y == 0 || other.Z == 0 {
		fmt.Println("Error: Division by zero in vector components")
		return &Vec3{} // Return a zero vector as fallback
	}
	return &Vec3{
		X: v.X / other.X,
		Y: v.Y / other.Y,
		Z: v.Z / other.Z,
	}
}

// Len3 calculates the length of a 3D vector
func Len3(v *Vec3) float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
}
