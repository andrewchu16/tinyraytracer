package processing

import (
	"tinyraytracer/geometry"
)

// Supersampling Antialiasing (Downsampling)

func SSA(bufp *[][]geometry.Vec3, samples int) *[][]geometry.Vec3 {
	var (
		newHeight = len(*bufp) / samples
		newWidth  = len((*bufp)[0]) / samples
	)

	// Initialize new buffer
	newBuf := make([][]geometry.Vec3, newHeight)

	for i := range newBuf {
		newBuf[i] = make([]geometry.Vec3, newWidth)
	}

	for y := range newHeight {
		for x := range newWidth {
			var color geometry.Vec3
			for dx := range samples {
				for dy := range samples {
					color.Add((*bufp)[y*samples+dy][x*samples+dx])
				}
			}
			newBuf[y][x] = geometry.Div(color, float64(samples*samples))
		}
	}

	return &newBuf
}
