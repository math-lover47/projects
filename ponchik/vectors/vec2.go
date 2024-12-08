package vectors

import (
	"fmt"
	"math"
)

type vec2 struct {
	x, y float32
}

// Initialize all components to the same value
func initValueVec2(val float32) *vec2 {
	return &vec2{
		x: val,
		y: val,
	}
}

func initValuesVec2(x_, y_ float32) *vec2 {
	return &vec2{
		x: x_,
		y: y_,
	}
}

// Add two vectors (returns a new vector)
func (v *vec2) Plus(other *vec2) *vec2 {
	return &vec2{
		x: v.x + other.x,
		y: v.y + other.y,
	}
}

// Subtract two vectors (returns a new vector)
func (v *vec2) Minus(other *vec2) *vec2 {
	return &vec2{
		x: v.x - other.x,
		y: v.y - other.y,
	}
}

// Multiply two vectors (returns a new vector)
func (v *vec2) Mult(other *vec2) *vec2 {
	return &vec2{
		x: v.x * other.x,
		y: v.y * other.y,
	}
}

// Divide two vectors (returns a new vector)
func (v *vec2) Div(other *vec2) *vec2 {
	// Check for division by zero
	if other.x == 0 || other.y == 0 {
		fmt.Println("Error: Division by zero in vector components")
		return &vec2{} // Return a zero vector as fallback
	}
	return &vec2{
		x: v.x / other.x,
		y: v.y / other.y,
	}
}

func Len2(v *vec2) float32 {
	return float32(math.Sqrt(float64(v.x*v.x + v.y*v.y)))
}
