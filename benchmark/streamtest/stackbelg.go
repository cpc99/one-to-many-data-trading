package streamtest

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
		fmt.Println(i, z)
	}
	return z
}

func Sqrt_new(x float64) float64 {
	pre, z := 0.0, 1.0
	for math.Abs(pre-z) > 1e-8 {
		pre = z
		z = (z + x/z) / 2
	}
	return z
}

func main() {
	num := 34555.0
	fmt.Println("self Sqrt result is:", Sqrt(num))
	fmt.Println("self Sqrt new result is:", Sqrt_new(num))
	fmt.Println("math Sqrt result is:", math.Sqrt(num))
}
