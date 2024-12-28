package geometry

import (
	"fmt"
	"math"
)

type Vec4 struct {
	W float64
	X float64
	Y float64
	Z float64
}

func NewVec4(w, x, y, z float64) Vec4 {
	return Vec4{w, x, y, z}
}

func (v Vec4) Add(v2 Vec4) Vec4 {
	return Vec4{v.W + v2.W, v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

func (v Vec4) Sub(v2 Vec4) Vec4 {
	return Vec4{v.W - v2.W, v.X - v2.X, v.Y - v2.Y, v.Z - v2.Z}
}

func (v Vec4) Mul(scalar float64) Vec4 {
	return Vec4{v.W * scalar, v.X * scalar, v.Y * scalar, v.Z * scalar}
}

func (v Vec4) Div(scalar float64) Vec4 {
	return Vec4{v.W / scalar, v.X / scalar, v.Y / scalar, v.Z / scalar}
}

func (v Vec4) Dot(v2 Vec4) float64 {
	return v.W*v2.W + v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v Vec4) Neg() Vec4 {
	return Vec4{-v.W, -v.X, -v.Y, -v.Z}
}

func (v Vec4) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vec4) Normalize() Vec4 {
	return v.Div(v.Length())
}

func (v Vec4) String() string {
	return fmt.Sprintf("(%f, %f, %f, %f)", v.W, v.X, v.Y, v.Z)
}
