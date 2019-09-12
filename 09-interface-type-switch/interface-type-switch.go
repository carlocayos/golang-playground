package main

import "fmt"

type Stringer interface {
	String() string
}

type StringerImpl struct {}

func (*StringerImpl) String() string {
	return "From Stringer"
}

func main() {


	var value interface{}
	// 1) Uncomment and output "Hello World"
	//value = "Hello World"
	// 2) Uncomment and output "From Stringer"
	value = &StringerImpl{}

	switch str := value.(type) {
	case string:
		fmt.Println(str)
	case Stringer:
		fmt.Println(str)
	}
}



