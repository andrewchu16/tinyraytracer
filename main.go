package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"tinyraytracer/geometry"
	"tinyraytracer/processeffects"
	"time"
)

const (
	WIDTH    = 1024
	HEIGHT   = 768
	IMG_NAME = "output.png"
)

func render(bufp *[][]geometry.Vec3) {
    buf := *bufp
	for y := range HEIGHT {
		for x := range WIDTH {
			r := math.Round(float64(x) / 64) * 64 / float64(WIDTH)
			g := 0.0
			b := math.Round(float64(y) / 64) * 64 / float64(HEIGHT)

            buf[y][x] = geometry.NewVec3(r, g, b)
		}
	}
}

func process(bufp *[][]geometry.Vec3) {
    *bufp = *processeffects.BoxBlur(bufp, 1)
}

func save(bufp *[][]geometry.Vec3) {
    buf := *bufp
	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))

	// Convert buffer to image
	for y := range HEIGHT {
		for x := range WIDTH {
			// Convert normalized color to 8-bit color
			red := uint8(buf[y][x].X * 255)
			green := uint8(buf[y][x].Y * 255)
			blue := uint8(buf[y][x].Z * 255)

			img.Set(x, y, color.RGBA{red, green, blue, 255}) // Alpha is always 255
		}
	}

	file, err := os.Create(IMG_NAME)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}
}

func timeIt(f func(), name string) {
	start := time.Now()
	fmt.Print(name, "...")
	f()
	fmt.Println(time.Since(start))
}

func main() {
	// Initialize buffer
	var buf [][]geometry.Vec3

	timeIt(func() { 
		buf = make([][]geometry.Vec3, HEIGHT)
		for i := range buf {
			buf[i] = make([]geometry.Vec3, WIDTH)
		}
	}, "Initialize buffer")

	// Render the scene
	timeIt(func() { render(&buf) }, "Rendering scene")

	// Post-process the image
	timeIt(func() { process(&buf) }, "Processing image")

	// Save the image
	timeIt(func() { save(&buf) }, "Saving image")

	fmt.Println("Done!")
}
