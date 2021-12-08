package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
func powElse(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println("Sqrt", sqrt(2), sqrt(-4))
	fmt.Println(
		"Pow",
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println(
		"powElse",
		powElse(3, 2, 10),
		powElse(3, 3, 20),
	)
}
