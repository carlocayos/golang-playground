package main

import "fmt"

func main() {
	fmt.Println("Appending to a slice")

	slice1 := []string{"a", "b", "c"}
	fmt.Println(slice1) // [a b c]

	slice2 := append(slice1, "d", "e")
	fmt.Println(slice2) // [a b c d e]

	slice3 := []string{"f", "g", "h", "i"}
	fmt.Println(slice3) // [g h i]

	slice4 := append(slice2, slice3...)
	fmt.Println(slice4) // [a b c d e g h i]
}
