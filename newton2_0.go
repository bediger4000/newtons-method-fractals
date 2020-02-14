package main

/* The Go Programming Language, exercis 3.7, page 62:
 * Another simple fractal use Newton's Method to find complex
 * solutions to a function such as z**4 - 1 = 0. Shade each
 * starting point by the number of iterations required to
 * get close to one of the four roots. Color each point by the
 * root it approaches.
 *
 * Roots of z**4 - 1 = 0 are 1, -1, i -i
 */

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2.0, -2.0, 2.0, 2.0
		width, height          = 4096, 4096
		contrast               = 20
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {

			if px == py {
				continue  // Newton's method for z**4 - 1 doesn't converge
			}

			x := float64(px)/height*(xmax-xmin) + xmin

			iters, vect := newtons(complex(x, y))

			r := real(vect)
			rabs := math.Abs(r)
			i := imag(vect)

			var red, green, blue uint8
			red, green, blue = 0, 0, 0

			shade := uint8(255 - contrast*iters)

			switch {
			case rabs < 0.1: // i or -i
				switch {
				case i > 0.8: // i
					red = shade
				case i < -0.8: // -1
					green = shade
				}

			case rabs > 0.9: // 1 or -1
				switch {
				case r > 0.9: // 1
					blue = shade
				case r < -0.9: // -1
					red = shade
					green = shade
				}
			}

			img.Set(px, py, color.RGBA{red, green, blue, 255})
		}
	}
	png.Encode(os.Stdout, img)
}

func newtons(z complex128) (int, complex128) {
	const iterations = 500
	var znext complex128
	var i int

	delta := 1.0 - cmplx.Abs(z)

	for i = 0; math.Abs(delta) > 0.01 && i < iterations; i++ {
		// f(z) = z^4 - 1
		// f/(z) = 3z^3
		// next z = z - (z^4 - 1)/4z^3)
		znext = (z - (z*z*z*z-1)/(4.0*z*z*z))
		z = znext
		delta = 1.0 - cmplx.Abs(znext)
	}

	return i, znext
}
