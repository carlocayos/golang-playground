package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func addPassFunc(f func(int, int) int) int {
	// add a fixed value
	return f(9, 20)
}

func main() {

	// invoke the function directly
	fmt.Println(add(1, 3))

	// pass the function
	fmt.Println(addPassFunc(add))
}
