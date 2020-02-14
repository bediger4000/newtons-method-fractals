package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"math"
	// "strconv"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -8.0, -8.0, 8.0, 8.0
		width, height          = 16384, 16384
		contrast               = 18
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {

			if px == py { continue }

			x := float64(px)/height*(xmax-xmin) + xmin

			iters, vect := newtons(complex(x, y))

			r := real(vect)

			var red, g uint8
			red, g = 0, 0

			shade := uint8(255 - contrast*iters)

			switch {
			case r > 0.0:
				red = shade
			case r < 0.0:
				g = shade
			}

			col := color.RGBA{red, g, 0, 255}

			img.Set(px, py, col)
		}
	}
/*
	for px := 0; px < width; px++ {
		img.Set(px, height/2, color.White)
	}
	for py := 0; py < height; py++ {
		img.Set(width/2, py, color.White)
	}
*/
	png.Encode(os.Stdout, img)
}

func newtons(z complex128) (int, complex128) {
	const iterations = 500
	var znext complex128
	var i int

    delta := 1.0 - cmplx.Abs(z)

	for i = 0; math.Abs(delta) > 0.01 && i < iterations; i++ {
		znext = (z - (z*z - 1)/(2.0*z))
		z = znext
		delta = 1.0 - cmplx.Abs(znext)
	}

	return i, znext
}
