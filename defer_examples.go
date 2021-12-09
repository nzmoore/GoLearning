package main

import "fmt"

func simpleDefer() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

func stackedDefer() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	simpleDefer()
	stackedDefer()
}
