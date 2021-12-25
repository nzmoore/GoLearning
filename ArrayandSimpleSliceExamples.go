package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slice includes first element up to but not including the last
	var s []int = primes[1:4]
	fmt.Println(s)

	// slices are references to an array
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	aa := names[0:2]
	bb := names[1:3]
	fmt.Println(aa, bb)

	bb[0] = "XXX"
	fmt.Println(aa, bb)
	fmt.Println(names)

}
