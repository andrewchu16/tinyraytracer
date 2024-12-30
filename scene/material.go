package scene

import (
	"fmt"
	"tinyraytracer/geometry"
)

type Material struct {
	RefractiveIndex  float64
	Albedo           geometry.Vec4
	DiffuseColor     geometry.Vec3
	SpecularExponent float64
}

func NewMaterial(refractiveIndex float64, albedo geometry.Vec4, diffuseColor geometry.Vec3, specularExponent float64) Material {
	return Material{refractiveIndex, albedo, diffuseColor, specularExponent}
}

func (m *Material) String() string {
	return fmt.Sprintf("Material{RefractiveIndex: %f, Albedo: %v, DiffuseColor: %v, SpecularExponent: %f}", m.RefractiveIndex, m.Albedo, m.DiffuseColor, m.SpecularExponent)
}
