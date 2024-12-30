package processeffects

import (
    "tinyraytracer/geometry"
    "sync"
)

func BoxBlur(bufp *[][]geometry.Color, blurWidth int) *[][]geometry.Color {
	if blurWidth <= 0 {
		return bufp
	}

    buf := *bufp

	height := len(buf)
	width := len((*bufp)[0])

    newBuf := make([][]geometry.Vec3, height)
    for i := range newBuf {
        newBuf[i] = make([]geometry.Vec3, width)
    }

    var wg sync.WaitGroup
    for y := 0; y < height; y++ {
        wg.Add(1)
        go func(y int) {
            defer wg.Done()
            for x := 0; x < width; x++ {
                newCol := geometry.NewVec3(0, 0, 0)
                count := 0
                for j := -blurWidth; j <= blurWidth; j++ {
                    for i := -blurWidth; i <= blurWidth; i++ {
                        nx := x + i
                        ny := y + j
                        if nx >= 0 && nx < width && ny >= 0 && ny < height {
                            newCol.Add(buf[ny][nx])
                            count++
                        }
                    }
                }
                newCol.Div(float64(count))
                newBuf[y][x] = newCol
            }
        }(y)
    }
    wg.Wait()

    return &newBuf
}