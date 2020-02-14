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
	iters, a := newtons(z)

	fmt.Printf("Answer: %v (%d)\n", a, iters)

}
func newtons(z complex128) (int, complex128) {
	const iterations = 750
	var znext complex128
	var i int

    delta := 1.0 - cmplx.Abs(z)

	for i = 0; math.Abs(delta) > 0.01 && i < iterations; i++ {
		znext = (z - (z*z*z*z*z*z - 1)/(6.0*z*z*z*z*z))
		z = znext
		delta = 1.0 - cmplx.Abs(znext)
	}

	return i, znext
}
