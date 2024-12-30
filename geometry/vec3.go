package geometry

import (
	"fmt"
	"math"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

var (
	ZERO_VEC3 = NewVec3(0, 0, 0)
	UNIT_X    = NewVec3(1, 0, 0)
	UNIT_Y    = NewVec3(0, 1, 0)
)

func NewVec3(x, y, z float64) Vec3 {
    return Vec3{x, y, z}
}

func (v *Vec3) Add(v2 Vec3) {
    v.X += v2.X
    v.Y += v2.Y
    v.Z += v2.Z
}

func Add(v1, v2 Vec3) Vec3 {
	return Vec3{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func (v *Vec3) Sub(v2 Vec3) {
    v.X -= v2.X
    v.Y -= v2.Y
    v.Z -= v2.Z
}

func Sub(v1, v2 Vec3) Vec3 {
	return Vec3{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

func (v *Vec3) Mul(scalar float64) {
    v.X *= scalar
    v.Y *= scalar
    v.Z *= scalar
}

func Mul(v1 Vec3, scalar float64) Vec3 {
	return Vec3{v1.X * scalar, v1.Y * scalar, v1.Z * scalar}
}

func (v *Vec3) Div(scalar float64) {
    v.X /= scalar
    v.Y /= scalar
    v.Z /= scalar
}

func Div(v1 Vec3, scalar float64) Vec3 {
	return Vec3{v1.X / scalar, v1.Y / scalar, v1.Z / scalar}
}

func (v *Vec3) Dot(v2 Vec3) float64 {
    return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func Dot(v1, v2 Vec3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func (v *Vec3) Negate() {
    v.X = -v.X
    v.Y = -v.Y
    v.Z = -v.Z
}

func (v Vec3) Neg() Vec3 {
	return Vec3{-v.X, -v.Y, -v.Z}
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.Dot(*v))
}

func Length(v Vec3) float64 {
	return math.Sqrt(v.Dot(v))
}

func (v *Vec3) SqrLength() float64 {
	return v.Dot(*v)
}

func (v *Vec3) Normalize() {
	v.Div(v.Length())
}

func (v Vec3) Normal() Vec3 {
	v.Normalize()

	return v
}

func (v *Vec3) IsNormalized() bool {
	return math.Abs(v.SqrLength()-1) < 0.00001
}

func (v *Vec3) Copy() Vec3 {
	return Vec3{v.X, v.Y, v.Z}
}

func (v Vec3) Cross(v2 Vec3) Vec3 {
	return Vec3{
		v.Y*v2.Z - v.Z*v2.Y,
		v.Z*v2.X - v.X*v2.Z,
		v.X*v2.Y - v.Y*v2.X,
	}
}

func Cross(v1, v2 Vec3) Vec3 {
	return Vec3{
		v1.Y*v2.Z - v1.Z*v2.Y,
		v1.Z*v2.X - v1.X*v2.Z,
		v1.X*v2.Y - v1.Y*v2.X,
	}
}

func (v Vec3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v.X, v.Y, v.Z)
}