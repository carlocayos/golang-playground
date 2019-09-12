package main

import "fmt"

func main() {

	x := [5]int{99,88,43,7,11}

	for _, val := range x {
		fmt.Printf("%T\t%d\n", val, val)
	}

	fmt.Printf("%T\n", x)
}
