package main

import "fmt"

func main() {
	x := []int{11,5,8,3,2,66,36,321,32,256}

	fmt.Println("length =", len(x))
	fmt.Println("cap =", cap(x))

	for _, val := range x {
		fmt.Printf("%T\t%d\n", val, val)
	}

	fmt.Printf("%T\n", x)
}
