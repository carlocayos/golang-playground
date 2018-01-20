//Interfaces are implemented implicitly
//A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.
//Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

package main

import "fmt"

// interface type with MyMethod()
type I interface {
	MyMethod()
}

// My struct h
type MyStruct struct {
	Name string
}

// implicit interface
func (m MyStruct) MyMethod() {
	fmt.Println("Entered MyMethod()")
}

func main() {
	// declare interface variable
	var myInterface I = MyStruct{"Carlo"}

	// call the method interface
	myInterface.MyMethod()
}
