package main

import (
	"fmt"
	"math"
	"time"
	"tinyraytracer/camera"
	"tinyraytracer/geometry"
	"tinyraytracer/processing"
)

const (
	WIDTH    = 1024
	HEIGHT   = 768
	IMG_NAME = "output.png"
)

func render(camera *camera.Camera) {
	for y := range HEIGHT {
		for x := range WIDTH {
			r := math.Round(float64(x) / 64) * 64 / float64(WIDTH)
			g := 0.0
			b := math.Round(float64(y) / 64) * 64 / float64(HEIGHT)

            camera.Buf[y][x] = geometry.NewVec3(r, g, b)
		}
	}
}

func process(camera *camera.Camera) {
    camera.Buf = *processing.BoxBlur(&camera.Buf, 1)
}

func timeIt(f func(), name string) {
	start := time.Now()
	fmt.Print(name, "...")
	f()
	fmt.Println(time.Since(start))
}

func main() {
	camera := camera.NewCamera(WIDTH, HEIGHT, IMG_NAME)
	
	// Initialize camera
	timeIt(camera.Init, "Initialize camera")

	// Render the scene
	timeIt(func() { render(camera) }, "Rendering scene")

	// Post-process the image
	timeIt(func() { process(camera) }, "Processing image")

	// Save the image
	timeIt(camera.Save, "Saving image")

	fmt.Println("Done!")
}
