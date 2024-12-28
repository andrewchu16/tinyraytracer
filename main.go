package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	. "tinyraytracer/geometry"
	"tinyraytracer/postprocess"
)

const (
	WIDTH    = 1024
	HEIGHT   = 768
	IMG_NAME = "output.png"
)

func render(bufp *[][]Vec3) {
    buf := *bufp
	for y := range HEIGHT {
		for x := range WIDTH {
			r := math.Round(float64(x) / 64) * 64 / float64(WIDTH)
			g := 0.0
			b := math.Round(float64(y) / 64) * 64 / float64(HEIGHT)

            buf[y][x] = NewVec3(r, g, b)
		}
	}
}

func process(bufp *[][]Vec3) {
    postprocess.BoxBlur(bufp, 5)
}

func save(bufp *[][]Vec3) {
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

func main() {
	buf := make([][]Vec3, HEIGHT)
	for i := range buf {
		buf[i] = make([]Vec3, WIDTH)
	}

	for y := range HEIGHT {
		for x := range WIDTH {
			buf[y][x] = NewVec3(0, 0, 0)
		}
	}

	// Render the scene
	fmt.Println("Rendering scene...")
	render(&buf)

	// Post-process the image
	fmt.Println("Post-processing image...")
	process(&buf)

	// Save the image
	fmt.Println("Saving image...")
	save(&buf)

	fmt.Println("Done!")
}
