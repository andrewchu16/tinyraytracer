package scene

import (
	"math"
	"tinyraytracer/geometry"
)

type Scene struct {
	Spheres  []*Sphere
	Lights   []*Light
	maxDepth int
}

func NewScene(maxDepth int) *Scene {
	return &Scene{
		Spheres:  make([]*Sphere, 0),
		Lights:   make([]*Light, 0),
		maxDepth: maxDepth,
	}
}

func (s *Scene) AddSphere(sphere *Sphere) {
	s.Spheres = append(s.Spheres, sphere)
}

func (s *Scene) AddLight(light *Light) {
	s.Lights = append(s.Lights, light)
}

func (s *Scene) CastRay(ray *Ray) geometry.Vec3 {
	return s.castRayDepth(ray, 0)
}

// Find the closest sphere that intersects with the ray
// Returns the normal at the intersection point, the intersection point itself, and the material of the sphere
func (s *Scene) findIntersect(ray *Ray) (bool, geometry.Vec3, geometry.Vec3, *Material) {
	closestSphereDist := math.MaxFloat64
	var normal geometry.Vec3
	var hitpoint geometry.Vec3
	var material *Material

	for _, sphere := range s.Spheres {
		intersects, dist := sphere.RayIntersect(ray)

		if intersects && dist < closestSphereDist {
			closestSphereDist = dist
			hitpoint = geometry.Add(ray.Orig, geometry.Mul(ray.Dir, dist))
			normal = geometry.Sub(hitpoint, sphere.Center).Normal()
			material = &sphere.Material
		}
	}

	return closestSphereDist < math.MaxFloat64, normal, hitpoint, material
}

func (s *Scene) castRayDepth(ray *Ray, depth int) geometry.Vec3 {
	intersects, normal, hitpoint, material := s.findIntersect(ray)

	// If the ray doesn't intersect any spheres, return the background color
	if depth > s.maxDepth || !intersects {
		return geometry.NewVec3(0.2, 0.7, 0.8)
	}

	var (
		reflectDir  = reflect(ray.Dir, normal).Normal()
		refractDir  = refract(ray.Dir, normal, material.RefractiveIndex).Normal()
		reflectOrig geometry.Vec3
		refractOrig geometry.Vec3
	)

	// Offset the hitpoint to avoid occlusion by the object itself
	if geometry.Dot(reflectDir, normal) < 0 {
		reflectOrig = geometry.Sub(hitpoint, geometry.Mul(normal, 0.001))
	} else {
		reflectOrig = geometry.Add(hitpoint, geometry.Mul(normal, 0.001))
	}

	if geometry.Dot(refractDir, normal) < 0 {
		refractOrig = geometry.Sub(hitpoint, geometry.Mul(normal, 0.001))
	} else {
		refractOrig = geometry.Add(hitpoint, geometry.Mul(normal, 0.001))
	}

	reflectColor := s.castRayDepth(&Ray{reflectOrig, reflectDir}, depth+1)
	refractColor := s.castRayDepth(&Ray{refractOrig, refractDir}, depth+1)

	var (
		diffuseLightIntensity  = 0.0
		specularLightIntensity = 0.0
	)

	for _, light := range s.Lights {
		lightDir := geometry.Sub(light.Position, hitpoint).Normal()
		lightDistance := geometry.Length(geometry.Sub(light.Position, hitpoint))

		var shadowOrig geometry.Vec3

		if geometry.Dot(lightDir, normal) < 0 {
			shadowOrig = geometry.Sub(hitpoint, geometry.Mul(normal, 0.001))
		} else {
			shadowOrig = geometry.Add(hitpoint, geometry.Mul(normal, 0.001))
		}

		shadowIntersects, _, shadowHitpoint, _ := s.findIntersect(&Ray{shadowOrig, lightDir})

		if shadowIntersects && geometry.Length(geometry.Sub(shadowHitpoint, shadowOrig)) < lightDistance {
			continue
		}

		diffuseLightIntensity += light.Intensity * math.Max(0, geometry.Dot(lightDir, normal))
		specularLightIntensity += math.Pow(math.Max(0, geometry.Dot(reflect(lightDir.Neg(), normal), ray.Dir.Neg())), material.SpecularExponent) * light.Intensity
	}

	var (
		diffuseColor  = geometry.Mul(material.DiffuseColor, diffuseLightIntensity*material.Albedo.W)
		specularColor = geometry.Mul(geometry.NewVec3(1, 1, 1), specularLightIntensity*material.Albedo.X)
	)
	reflectColor = geometry.Mul(reflectColor, material.Albedo.Y)
	refractColor = geometry.Mul(refractColor, material.Albedo.Z)

	return geometry.Add(geometry.Add(diffuseColor, specularColor), geometry.Add(reflectColor, refractColor))
}

func reflect(v, normal geometry.Vec3) geometry.Vec3 {
	return geometry.Sub(v, geometry.Mul(normal, 2*geometry.Dot(v, normal)))
}

func refract(v, normal geometry.Vec3, refractiveIndex float64) geometry.Vec3 {
	var (
		cosI = -math.Max(-1, math.Min(1, geometry.Dot(v, normal)))
		etaI = 1.0
		etaT = refractiveIndex
	)

	if cosI < 0 {
		cosI = -cosI
		etaI, etaT = etaT, etaI
		normal.Negate()
	}

	eta := etaI / etaT
	k := 1 - eta*eta*(1-cosI*cosI)

	if k < 0 {
		return geometry.ZERO_VEC3
	}

	return geometry.Add(geometry.Mul(v, eta), geometry.Mul(normal, eta*cosI-math.Sqrt(k)))
}
