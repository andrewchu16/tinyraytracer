package camera

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"sync"
	"tinyraytracer/geometry"
	"tinyraytracer/scene"
)

type Camera struct {
	width   int
	height  int
	fov     int
	imgName string
	Buf     [][]geometry.Vec3
	scene   *scene.Scene
}

func NewCamera(width, height int, fov int, imgName string) *Camera {
	return &Camera{
		width:   width,
		height:  height,
		fov:     fov,
		imgName: imgName,
		Buf:     nil,
		scene:   nil,
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

func (c *Camera) Render() {
	var wg sync.WaitGroup
	for j := range c.height {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()

			for i := range c.width {
				x := (2.0*(float64(i)+0.5)/float64(c.width) - 1.0) * math.Tan(float64(c.fov)/2.0) * float64(c.width) / float64(c.height)
				y := -(2.0*(float64(j)+0.5)/float64(c.height) - 1.0) * math.Tan(float64(c.fov)/2.0)

				orig := geometry.ZERO_VEC3.Copy()
				dir := geometry.NewVec3(x, y, -1).Normal()
				ray := scene.NewRay(&orig, &dir)

				color := c.scene.CastRay(&ray)

				// Normalize color to 1
				maxColorChannel := math.Max(color.X, math.Max(color.Y, color.Z))

				if maxColorChannel > 1 {
					color.Div(maxColorChannel)

					// color.X = math.Min(1, color.X)
					// color.Y = math.Min(1, color.Y)
					// color.Z = math.Min(1, color.Z)
				}

				c.Buf[j][i] = color
			}
		}(j)
	}

	wg.Wait()
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
