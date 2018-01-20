package main

import (
	"fmt"
)

func main() {

	// Arrays - with size
	names := [3]string{"John", "Paul", "Greg"}
	fmt.Println(names)

	// Arrays - have the compiler count the array elements for you
	cars := [...]string{"BMW", "Toyota", "Honda", "Mazda"}
	fmt.Println(cars)

	// Slice - without size
	food := []string{"pizza", "ice cream", "burgers"}
	fmt.Println(food)

	// show slice size and cap
	// size = 3, cap = 3, [pizza ice cream burgers]
	fmt.Printf("size = %d, cap = %d, %v", len(food), cap(food), food)

}
