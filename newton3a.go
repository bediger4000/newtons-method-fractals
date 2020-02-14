package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
	// "strconv"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -1.0, -1.0, 1.0, 1.0
		width, height          = 1000, 1000
		contrast               = 19
	)
	/*
		xmin, _ := strconv.ParseFloat(os.Args[1], 64)
		ymin, _ := strconv.ParseFloat(os.Args[2], 64)
		xmax, _ := strconv.ParseFloat(os.Args[3], 64)
		ymax, _ := strconv.ParseFloat(os.Args[4], 64)
	*/
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {

			if px == py {
				continue
			}

			x := float64(px)/height*(xmax-xmin) + xmin

			iters, vect := newtons(complex(x, y))

			i := imag(vect)
			r := real(vect)

			var red, g, b uint8
			red, g, b = 0, 0, 0

			shade := uint8(255 - contrast*iters)

			// (0.5+0.8660254037i)
			// (0.5-0.8660254037i)
			switch {
			case r > 0.6:
				b = shade
			case i > 0.0:
				red = shade
			case i < 0.0:
				g = shade
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
	const iterations = 500
	var znext complex128
	var i int

	delta := 1.0 - cmplx.Abs(z)

	for i = 0; math.Abs(delta) > 0.01 && i < iterations; i++ {
		znext = (z - (z*z*z-1)/(3.0*z*z))
		z = znext
		delta = 1.0 - cmplx.Abs(znext)
	}

	return i, znext
}
