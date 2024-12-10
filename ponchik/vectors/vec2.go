package vectors

import (
	"fmt"
	"math"
)

type Vec2 struct {
	X, Y float64
}

// InitValueVec2 initializes all components to the same value
func InitValueVec2(val float64) *Vec2 {
	return &Vec2{
		X: val,
		Y: val,
	}
}

// InitValuesVec2 initializes vector with specific X and Y values
func InitValuesVec2(X, Y float64) *Vec2 {
	return &Vec2{
		X: X,
		Y: Y,
	}
}

// Plus adds two vectors (returns a new vector)
func (v *Vec2) Plus(other *Vec2) *Vec2 {
	return &Vec2{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

// Minus subtracts two vectors (returns a new vector)
func (v *Vec2) Minus(other *Vec2) *Vec2 {
	return &Vec2{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

// Mult multiplies two vectors (returns a new vector)
func (v *Vec2) Mult(other *Vec2) *Vec2 {
	return &Vec2{
		X: v.X * other.X,
		Y: v.Y * other.Y,
	}
}

// Div divides two vectors (returns a new vector)
func (v *Vec2) Div(other *Vec2) *Vec2 {
	// Check for division by zero
	if other.X == 0 || other.Y == 0 {
		fmt.Println("Error: Division by zero in vector components")
		return &Vec2{} // Return a zero vector as fallback
	}
	return &Vec2{
		X: v.X / other.X,
		Y: v.Y / other.Y,
	}
}

// Len2 calculates the length of a 2D vector
func Len2(v *Vec2) float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}
