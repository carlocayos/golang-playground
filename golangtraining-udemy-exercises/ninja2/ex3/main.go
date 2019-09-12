package main

import "fmt"

const (
	a int    = 5
	b        = 655
	c string = "Hello world"
	d        = "Hello untyped"
)

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	fmt.Println(d)
	fmt.Printf("%T\n", d)
}
