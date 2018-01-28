package main

import "fmt"

func main() {

	defer fmt.Println("DEFER CALL")
	defer fmt.Println("DEFER CALL2")
	defer fmt.Println("DEFER CALL3")
	fmt.Println("After the defer code line")

	//After the defer code line
	//DEFER CALL3
	//DEFER CALL2
	//DEFER CALL
}


