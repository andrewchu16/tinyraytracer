package main

import (
	"fmt"
	"math"
	"time"
	"tinyraytracer/camera"
	"tinyraytracer/geometry"
	"tinyraytracer/processing"
	"tinyraytracer/scene"
)

const (
	WIDTH    = 1024 * 4
	HEIGHT   = 768 * 4
	IMG_NAME = "output.png"
)

func process(camera *camera.Camera) {
	// camera.Buf = *processing.BoxBlur(&camera.Buf, 1)
	camera.SetBuf(*processing.SSA(&camera.Buf, 4))
}

func timeIt(f func(), name string) {
	start := time.Now()
	fmt.Print(name, "...")
	f()
	fmt.Println(time.Since(start))
}

func main() {
	camera := camera.NewCamera(WIDTH, HEIGHT, int(math.Round(math.Pi/3.0)), IMG_NAME)

	sc := scene.NewScene(4)

	// Initialize scene
	timeIt(func() {
		// Materials
		var (
			ivory = scene.NewMaterial(1.0, geometry.NewVec4(0.6, 0.3, 0.1, 0.0), geometry.NewVec3(0.4, 0.4, 0.3), 50.0)
			glass = scene.NewMaterial(1.5, geometry.NewVec4(0.0, 0.5, 0.1, 0.8), geometry.NewVec3(0.6, 0.7, 0.8), 125.0)
			redRubber = scene.NewMaterial(1.0, geometry.NewVec4(0.9, 0.1, 0.0, 0.0), geometry.NewVec3(0.3, 0.1, 0.1), 10.0)
			mirror = scene.NewMaterial(1.0, geometry.NewVec4(0.0, 10.0, 0.8, 0.0), geometry.NewVec3(1.0, 1.0, 1.0), 1425.0)
		)

		// Spheres
		var (
			sphere1 = scene.NewSphere(geometry.NewVec3(-3, 0, -16), 2, &ivory)
			sphere2 = scene.NewSphere(geometry.NewVec3(-1.0, -1.5, -12), 2, &glass)
			sphere3 = scene.NewSphere(geometry.NewVec3(1.5, -0.5, -18), 3, &redRubber)
			sphere4 = scene.NewSphere(geometry.NewVec3(7, 5, -18), 4, &mirror)
		)

		// Add spheres to scene
		sc.AddSphere(&sphere1)
		sc.AddSphere(&sphere2)
		sc.AddSphere(&sphere3)
		sc.AddSphere(&sphere4)

		// Lights
		var (
			light1 = scene.NewLight(geometry.NewVec3(-20, 20, 20), 1.5)
			light2 = scene.NewLight(geometry.NewVec3(30, 50, -25), 1.8)
			light3 = scene.NewLight(geometry.NewVec3(30, 20, 30), 1.7)
		)

		// Add lights to scene
		sc.AddLight(&light1)
		sc.AddLight(&light2)
		sc.AddLight(&light3)
	}, "Initialize scene")

	// Initialize camera
	timeIt(func() {
		camera.SetScene(sc)
		camera.Init()
	}, "Initialize camera")

	// Render the scene
	timeIt(camera.Render, "Rendering scene")

	// Post-process the image
	timeIt(func() { process(camera) }, "Processing image")

	// Save the image
	timeIt(camera.Save, "Saving image")

	fmt.Println("Done!")
}
