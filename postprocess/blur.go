package postprocess

import . "tinyraytracer/geometry"

func BoxBlur(bufp *[][]Vec3, blurWidth int) {
	if blurWidth <= 0 {
		return
	}

    buf := *bufp

	height := len(buf)
	width := len((*bufp)[0])

    copyBuf := make([][]Vec3, height)
    for i := range copyBuf {
        copyBuf[i] = make([]Vec3, width)
    }

    for y := range height {
        for x := range width {
            // Blur the image
            r := 0.0
            g := 0.0
            b := 0.0

            count := 0
            for j := -blurWidth; j <= blurWidth; j++ {
                for i := -blurWidth; i <= blurWidth; i++ {
                    nx := x + i
                    ny := y + j

                    if nx >= 0 && nx < width && ny >= 0 && ny < height {
                        r += buf[ny][nx].X
                        g += buf[ny][nx].Y
                        b += buf[ny][nx].Z
                        count++
                    }
                }
            }

            r /= float64(count)
            g /= float64(count)
            b /= float64(count)

            copyBuf[y][x] = NewVec3(r, g, b)
        }
    }

    *bufp = copyBuf
}