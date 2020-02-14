package main
import (
/*
	"image"
	"image/color"
	"image/png"
*/
	"math"
	"math/cmplx"
	"os"
	"strconv"
	"fmt"
)

func main() {
	x, _ := strconv.ParseFloat(os.Args[1], 64)
	y, _ := strconv.ParseFloat(os.Args[2], 64)

	z := complex(x,y)
	delta := 1.0 - cmplx.Abs(z)
	fmt.Printf("(x,y) = (%v, %v)\n", x, y)
	fmt.Printf("delta = %v\n", delta)

	for i:=0; math.Abs(delta) > 0.01 && i < 2000; i++ {
		znext := (z - (z*z*z*z - 1)/(4.0*z*z*z))
		delta = 1.0 - cmplx.Abs(znext)
		fmt.Printf("%v\t%v\t%v\n", real(znext), imag(znext), delta)
		z = znext
	}
}
