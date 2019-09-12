package main

import "fmt"

type donut int

var x donut
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)

	// conversion of x into its underlying type
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
