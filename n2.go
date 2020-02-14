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
	"fmt"
)

func main() {

	for x := -2.0; x < 2.0; x += 0.01 {
		z := complex(x,-0.35)
		_, a := newtons(z)
		fmt.Printf("%v\n", a)
	}

	os.Exit(0)
}
func newtons(z complex128) (int, complex128) {
	const iterations = 750
	var znext complex128
	var i int

    delta := 1.0 - cmplx.Abs(z)

	for i = 0; math.Abs(delta) > 0.01 && i < iterations; i++ {
		znext = (z - (z*z*z - 1)/(3.0*z*z))
		z = znext
		delta = 1.0 - cmplx.Abs(znext)
	}

	return i, znext
}
