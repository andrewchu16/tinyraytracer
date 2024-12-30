package scene

import (
	"fmt"
	"tinyraytracer/geometry"
)

type Light struct {
	Position  geometry.Vec3
	Intensity float64
}

func NewLight(position geometry.Vec3, intensity float64) Light {
	return Light{position, intensity}
}

func (l *Light) String() string {
	return fmt.Sprintf("Light{position: %v, intensity: %v}", l.Position, l.Intensity)
}
