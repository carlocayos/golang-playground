package main

import (
	"fmt"
)

type Dog struct {
	Name string
	Age int
}


func (d Dog) makeSound() string {
	return "Arf"
}

func main() {

	d := Dog{"bea", 5}
	fmt.Println(d.makeSound())
}
