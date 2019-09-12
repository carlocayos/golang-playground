package main

import "fmt"

func main() {
	fmt.Println("Switch Test")

	a := &Animal{}
	a.Eat()
	a.Sound()

	i := 1
	switch i {
	default:
		fmt.Println("This is the default")
	case 1:
		fmt.Println("Case 1")
	}
}

type Animal struct {}

func (*Animal) Eat() {
	fmt.Println("Animal is eating")
}

func (*Animal) Sound() {
	fmt.Println("This is the animal sound")
}
