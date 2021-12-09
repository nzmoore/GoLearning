package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	z := float64(1)
	if x < 0 {
		return 0, errors.New("Cannot handle negative numbers ")
	}
	if x == 0 {
		return x, nil
	} else {

		latest := float64(999999999)
		for i := 0; i < 10000; i++ {
			z -= (z*z - x) / (2 * z)
			if math.Abs(latest-z) < 0.0000000001 {
				return z, nil
			} else {
				latest = z
				println(i, z, latest)
			}
		}
	}
	return z, nil
}

func main() {

	fmt.Println(Sqrt(0))
	fmt.Println(Sqrt(9))
	fmt.Println(Sqrt(-9))

}
