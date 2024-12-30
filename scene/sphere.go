package scene

import (
	"fmt"
	"math"
	"tinyraytracer/geometry"
)

type Sphere struct {
	Center   geometry.Vec3
	Radius   float64
	Material Material
}

func NewSphere(center geometry.Vec3, radius float64, material *Material) Sphere {
	return Sphere{center, radius, *material}
}

// RayIntersect returns true if the ray intersects the sphere, and the scalar of the ray at the point of intersection
// If the ray does not intersect the sphere, the scalar is NaN
func (s *Sphere) RayIntersect(ray *Ray) (bool, float64) {
	L := geometry.Sub(s.Center, ray.Orig)
	tca := geometry.Dot(L, ray.Dir)
	d2 := geometry.Dot(L, L) - tca*tca

	// If the ray doesn't intersect the sphere, return false
	if d2 > s.Radius*s.Radius {
		return false, math.NaN()
	}

	thc := math.Sqrt(s.Radius*s.Radius - d2)

	// Find the closest point of intersection
	t0 := tca - thc
	t1 := tca + thc

	if t0 < 0 {
		t0 = t1
	}

	// If the closest point of intersection is behind the ray's origin, return false
	if t0 < 0 {
		return false, math.NaN()
	}

	return true, t0
}

func (s *Sphere) String() string {
	return fmt.Sprintf("Sphere{Center: %v, Radius: %f, Material: %v}", s.Center, s.Radius, s.Material)
}
