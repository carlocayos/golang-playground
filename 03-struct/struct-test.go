package main

import (
	"fmt"
)

type Dog struct {
	Breed, Name string
	Age int
}

func main() {
	beaDog := Dog{Age: 1}
	beaDog.Breed = "Dachshund"
	beaDog.Name = "Bea"

	fmt.Println(beaDog)
}
