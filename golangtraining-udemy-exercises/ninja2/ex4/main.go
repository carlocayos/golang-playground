package main

import "fmt"

func main() {
	x := 100

	fmt.Printf("%v\t%b\t\t%#x\n", x, x ,x )

	y := x << 1
	fmt.Printf("%v\t%b\t%#x", y, y ,y )

}
