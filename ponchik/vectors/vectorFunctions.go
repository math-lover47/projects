package vectors

import (
	"math"
)

func Clamp(val, min, max float64) float32 {
	return float32(math.Max(math.Min(val, max), min))
}

func SignVal(n float32) float32 {
	if n > 0 {
		return 1
	}
	if n < 0 {
		return -1
	}
	return 0
}

func StepVal(edge, x float32) float32 {
	if x > edge {
		return 1
	}
	return 0
}

func Norm(v *vec3) *vec3 {
	return initValuesVec3(
		v.x/Len3(v),
		v.y/Len3(v),
		v.z/Len3(v),
	)
}

func Dot(v1, v2 *vec3) float32 {
	return v1.x*v2.x + v1.y*v2.y + v1.z*v2.z
}

func Abs(v *vec3) *vec3 {
	return initValuesVec3(
		float32(math.Abs(float64(v.x))),
		float32(math.Abs(float64(v.y))),
		float32(math.Abs(float64(v.z))),
	)
}

func SignVec(v *vec3) *vec3 {
	return initValuesVec3(
		SignVal(v.x),
		SignVal(v.y),
		SignVal(v.z),
	)
}

func StepVec(edge, v *vec3) *vec3 {
	return initValuesVec3(StepVal(edge.x, v.x), StepVal(edge.y, v.y), StepVal(edge.z, v.z))
}

func Reflect(rayDir, v *vec3) *vec3 {
	v.Mult(initValueVec3(2 * Dot(v, rayDir)))
	return rayDir.Minus(v)
}

func RotateX(v1 *vec3, angle float64) *vec3 {
	v2 := initValuesVec3(v1.x, v1.y, v1.z)
	v2.z = v1.z*float32(math.Cos(angle)) - v1.y*float32(math.Sin(angle))
	v2.y = v1.z*float32(math.Sin(angle)) + v1.y*float32(math.Cos(angle))
	return v2
}

func RotateY(v1 *vec3, angle float64) *vec3 {
	v2 := initValuesVec3(v1.x, v1.y, v1.z)
	v2.x = v1.x*float32(math.Cos(angle)) - v1.z*float32(math.Sin(angle))
	v2.z = v1.x*float32(math.Sin(angle)) + v1.z*float32(math.Cos(angle))
	return v2
}

func RotateZ(v1 *vec3, angle float64) *vec3 {
	v2 := initValuesVec3(v1.x, v1.y, v1.z)
	v2.x = v1.x*float32(math.Cos(angle)) - v1.y*float32(math.Sin(angle))
	v2.y = v1.x*float32(math.Sin(angle)) + v1.y*float32(math.Cos(angle))
	return v2
}

func Sphere(rayOrigin, rayDir *vec3, radius float32) *vec2 {
	b := Dot(rayOrigin, rayDir)
	c := Dot(rayOrigin, rayOrigin) - radius*radius
	h := b*b - c
	if h < 0 {
		return initValueVec2(-1)
	}
	h = float32(math.Sqrt(float64(h)))
	return initValuesVec2(-b-h, -b+h)
}

func Box(rayOrigin, rayDir, boxSize *vec3, outNormal **vec3) *vec2 {
	m := initValueVec3(1).Div(rayDir)
	n := m.Mult(rayOrigin)
	k := Abs(m).Mult(boxSize)
	t1 := n.OppositeSign().Minus(k)
	t2 := n.OppositeSign().Plus(k)
	tN := Max32(Max32(t1.x, t1.y), t1.z)
	tF := Min32(Min32(t2.x, t2.y), t2.z)

	if tN > tF || tF < 0 {
		return initValueVec2(-1)
	}
	yzx := initValuesVec3(t1.y, t1.z, t1.x)
	zxy := initValuesVec3(t1.z, t1.x, t1.y)
	*outNormal = SignVec(rayDir).OppositeSign().Mult(StepVec(yzx, t1)).Mult(StepVec(zxy, t1))
	return initValuesVec2(tN, tF)
}

func Plane(rayOrigin, rayDir, p *vec3, w float32) float32 {
	return -(Dot(rayOrigin, p) + w) / Dot(rayDir, p)
}

func Max32(a, b float32) float32 {
	return float32(math.Max(float64(a), float64(b)))
}

func Min32(a, b float32) float32 {
	return float32(math.Min(float64(a), float64(b)))
}
