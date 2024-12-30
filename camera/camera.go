package camera

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"tinyraytracer/geometry"
	"tinyraytracer/scene"
)

type Camera struct {
	width int
	height int
	imgName string
	Buf [][]geometry.Vec3
	scene *scene.Scene
}

func NewCamera(width, height int, imgName string) *Camera {
	return &Camera{
		width: width,
		height: height,
		imgName: imgName,
		Buf: nil,
		scene: nil,
	}
}

func (c *Camera) Init() {
	c.Buf = make([][]geometry.Vec3, c.height)
	for y := range c.Buf {
		c.Buf[y] = make([]geometry.Vec3, c.width)
	}
}

func (c *Camera) SetScene(scene *scene.Scene) {
	c.scene = scene
}

func (c *Camera) Save() {
	img := image.NewRGBA(image.Rect(0, 0, c.width, c.height))

	// Convert buffer to image
	for y := range c.height {
		for x := range c.width {
			// Convert normalized color to 8-bit color
			red := uint8(c.Buf[y][x].X * 255)
			green := uint8(c.Buf[y][x].Y * 255)
			blue := uint8(c.Buf[y][x].Z * 255)

			img.Set(x, y, color.RGBA{red, green, blue, 255}) // Alpha is always 255
		}
	}

	file, err := os.Create(c.imgName)
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
