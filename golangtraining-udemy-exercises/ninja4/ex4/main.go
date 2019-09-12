package main

import "fmt"

func main() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	result := append(x, 52)
	fmt.Println(result)

	result = append(x, 53, 54, 55)
	fmt.Println(result)

	y := []int{56, 57, 58, 59, 60}
	result = append(x, y...)
	fmt.Println(result)
}
