package main

import "fmt"

type I interface {
	MyMethod() string
}

// Dog struct
type Dog struct {
	Name string
}
func (d Dog) MyMethod() string {
	return "Hello World"
}


// Person struct
type Person struct {
	LastName string
}
func (p *Person) MyMethod() string {
	return "I am a person"
}

func main() {
	var i I = Dog{"bea"}
	fmt.Println(i.MyMethod()) // Hello World

	i = &Person{"My Last Name"}
	fmt.Println(i.MyMethod()) // I am a person

	person := Person{"Cayos"}
	fmt.Println(person.MyMethod()) // I am a person
}
