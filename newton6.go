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
		xmin, ymin, xmax, ymax = 0.0, 0.0, 2.0, 2.0
		width, height          = 8192, 8192
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
			i := imag(vect)

			var red, g, b uint8

			shade := uint8(255 - contrast*iters)

			// +-0.50 +- 0.86i
			// +- 1
			switch {
			case r < 0 && r < -0.60:   // -1
				red = shade
			case r > 0 && r > 0.49:   // 1
				g = shade
			case r > 0 && r < 0.51:
				switch {
				case i < 0:
					red = shade
					g = shade
				case i > 0:
					red = shade
					b = shade
				}
			case r < 0 && r < -0.4:
				switch {
				case i < 0:
					g = shade
					b = shade
				case i > 0:
					b = shade
				}
			}

			col := color.RGBA{red, g, b, 255}

			img.Set(px, py, col)
		}
	}
	for px := 0; px < width; px++ {
		img.Set(px, height/2, color.White)
	}
	for py := 0; py < height; py++ {
		img.Set(width/2, py, color.White)
	}
	png.Encode(os.Stdout, img)
}

func newtons(z complex128) (int, complex128) {
	const iterations = 750
	var znext complex128
	var i int

    delta := 1.0 - cmplx.Abs(z)

	for i = 0; math.Abs(delta) > 0.001 && i < iterations; i++ {
		znext = (z - (z*z*z*z*z*z - 1)/(6.0*z*z*z*z*z))
		z = znext
		delta = 1.0 - cmplx.Abs(znext)
	}

	return i, znext
}
