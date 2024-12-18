package vectors

import (
	"math"
)

func Clamp(val, min, max float64) float64 {
	return math.Max(math.Min(val, max), min)
}

func SignVal(n float64) float64 {
	if n > 0 {
		return 1
	}
	if n < 0 {
		return -1
	}
	return 0
}

func StepVal(edge, x float64) float64 {
	if x > edge {
		return 1
	}
	return 0
}

func Norm(v *Vec3) *Vec3 {
	length := Len3(v)
	if length == 0 {
		return InitValueVec3(0)
	}
	return InitValuesVec3(
		v.X/length,
		v.Y/length,
		v.Z/length,
	)
}

func Dot(v1, v2 *Vec3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func Abs(v *Vec3) *Vec3 {
	return InitValuesVec3(
		math.Abs(v.X),
		math.Abs(v.Y),
		math.Abs(v.Z),
	)
}

func SignVec(v *Vec3) *Vec3 {
	return InitValuesVec3(
		SignVal(v.X),
		SignVal(v.Y),
		SignVal(v.Z),
	)
}

func StepVec(edge, v *Vec3) *Vec3 {
	return InitValuesVec3(
		StepVal(edge.X, v.X),
		StepVal(edge.Y, v.Y),
		StepVal(edge.Z, v.Z),
	)
}

func Reflect(rayDir, normal *Vec3) *Vec3 {
	dotProduct := 2 * Dot(normal, rayDir)
	return InitValuesVec3(
		rayDir.X-dotProduct*normal.X,
		rayDir.Y-dotProduct*normal.Y,
		rayDir.Z-dotProduct*normal.Z,
	)
}

func RotateX(v1 *Vec3, angle float64) *Vec3 {
	v2 := InitValuesVec3(v1.X, v1.Y, v1.Z)
	v2.Z = v1.Z*math.Cos(angle) - v1.Y*math.Sin(angle)
	v2.Y = v1.Z*math.Sin(angle) + v1.Y*math.Cos(angle)
	return v2
}

func RotateY(v1 *Vec3, angle float64) *Vec3 {
	v2 := InitValuesVec3(v1.X, v1.Y, v1.Z)
	v2.X = v1.X*math.Cos(angle) - v1.Z*math.Sin(angle)
	v2.Z = v1.X*math.Sin(angle) + v1.Z*math.Cos(angle)
	return v2
}

func RotateZ(v1 *Vec3, angle float64) *Vec3 {
	v2 := InitValuesVec3(v1.X, v1.Y, v1.Z)
	v2.X = v1.X*math.Cos(angle) - v1.Y*math.Sin(angle)
	v2.Y = v1.X*math.Sin(angle) + v1.Y*math.Cos(angle)
	return v2
}

func Sphere(rayOrigin, rayDir *Vec3, radius float64) *Vec2 {
	b := Dot(rayOrigin, rayDir)
	c := Dot(rayOrigin, rayOrigin) - radius*radius
	h := b*b - c
	if h < 0.0 {
		return InitValueVec2(-1)
	}
	h = math.Sqrt(float64(h))
	return InitValuesVec2(-b-h, -b+h)
}

func Box(rayOrigin, rayDir, boxSize *Vec3, outNormal **Vec3) *Vec2 {
	m := InitValueVec3(1).Div(rayDir)
	n := m.Mult(rayOrigin)
	k := Abs(m).Mult(boxSize)
	t1 := n.OppositeSign().Minus(k)
	t2 := n.OppositeSign().Plus(k)
	tN := math.Max(math.Max(t1.X, t1.Y), t1.Z)
	tF := math.Min(math.Min(t2.X, t2.Y), t2.Z)
	if tN > tF || tF < 0 {
		return InitValueVec2(-1)
	}
	yzx := InitValuesVec3(t1.Y, t1.Z, t1.X)
	zxy := InitValuesVec3(t1.Z, t1.X, t1.Y)
	*outNormal = SignVec(rayDir).OppositeSign().Mult(StepVec(yzx, t1)).Mult(StepVec(zxy, t1))
	return InitValuesVec2(tN, tF)
}

func Plane(rayOrigin, rayDir, p *Vec3, w float64) float64 {
	return -((Dot(rayOrigin, p) + w) / Dot(rayDir, p))
}

func GetDist(p *Vec3, t float64) float64 {
	q := InitValuesVec2(Len2(InitValuesVec2(p.X, p.Y))-1.0, p.Z)
	return Len2(q) - 0.5
}
