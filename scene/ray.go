package scene

import (
	"fmt"
	"tinyraytracer/geometry"
)

type Ray struct {
	Orig geometry.Vec3
	Dir geometry.Vec3
}

func NewRay(orig, dir *geometry.Vec3) Ray {
	return Ray{*orig, *dir}
}

func (r *Ray) String() string {
	return fmt.Sprintf("Ray{orig: %v, dir: %v}", r.Orig, r.Dir)
}