package main

import "fmt"

func main() {

	var i interface{} = 1
	fmt.Println(i)

	result := i.(int)
	fmt.Println(result)

	result, ok := i.(int)
	fmt.Println(result, ok)

	result2, ok := i.(float64) // if type does not match then T will hold the zero value
	fmt.Println(result2, ok)

	result2 = i.(float64) // panic panic: interface conversion: interface {} is int, not float64
	fmt.Println(result2)
}
