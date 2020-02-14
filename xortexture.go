package main
/*
 * See: http://lodev.org/cgtutor/xortexture.html
 */

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	const (
		width, height          = 8192, 8192
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			c := uint8(px ^ py)
			col := color.RGBA{c, c, c, 255}
			img.Set(px, py, col)
		}
	}
	png.Encode(os.Stdout, img)
}
