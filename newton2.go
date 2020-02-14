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
		xmin, ymin, xmax, ymax = -2.0, -2.0, 2.0, 2.0
		width, height          = 8192, 8192
		contrast               = 18
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

			if px == py { continue }

			x := float64(px)/height*(xmax-xmin) + xmin

			iters, vect := newtons(complex(x, y))

			r := real(vect)
			rabs := math.Abs(r)
			i := imag(vect)

			var red, g, b uint8
			red, g, b = 0, 0, 0

			shade := uint8(255 - contrast*iters)

			switch {
				case  rabs < 0.1:  // i or -i
				switch {
					case i > 0.8:   // i
						red = shade
					case i < -0.8:  // -1
					g = shade
				}

				case  rabs > 0.9:  // 1 or -1
				switch {
					case r > 0.9:  // 1
						b = shade
					case r < -0.9:  // -1
						red = shade
						g = shade
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
	const iterations = 500
	var znext complex128
	var i int

    delta := 1.0 - cmplx.Abs(z)

	for i = 0; math.Abs(delta) > 0.01 && i < iterations; i++ {
		znext = (z - (z*z*z*z - 1)/(4.0*z*z*z))
		z = znext
		delta = 1.0 - cmplx.Abs(znext)
	}

	return i, znext
}
