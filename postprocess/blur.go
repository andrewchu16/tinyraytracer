package postprocess

import . "tinyraytracer/geometry"

func BoxBlur(bufp *[][]Vec3, blurWidth int) {
	if blurWidth <= 0 {
		return
	}

    buf := *bufp

	height := len(buf)
	width := len((*bufp)[0])

    newBuf := make([][]Vec3, height)
    for i := range newBuf {
        newBuf[i] = make([]Vec3, width)
    }

    for y := range height {
        for x := range width {
            // Blur the image
            newCol := NewVec3(0, 0, 0)

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
    }

    *bufp = newBuf
}