package main

import (
	"fmt"
)

// struct anonymous field - http://golangtutorials.blogspot.com.au/2011/06/anonymous-fields-in-structs-like-object.html
func main() {

	m := Mailman{
		Person{Name: "Mr Orange", LastName: "Fields"},
	}

	fmt.Printf("%+v\n", m)
	fmt.Printf("%v\n", m)

	// Mailman has an anonymous field Person

	//type Person struct {
	//	Name, LastName string
	//}
	//
	//type Mailman struct {
	//	Person
	//}


	// anonymous field exported values can be accessed directly
	fmt.Println(m.Name, m.LastName)
	// or through the defautl field name
	fmt.Println(m.Person.Name, m.Person.LastName)
}
