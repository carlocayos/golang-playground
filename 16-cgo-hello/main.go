package main

/*
 #include "hello.c"
 #include "sum.c"
*/
import "C"
import "fmt"

func main() {
	//===========================
	// say hello world - hello.c
	//===========================
	C.Hello()

	//===========================
	// add sum - sum.c
	//===========================
	a, b := 7, 4

	// convert to C integer
	aCint := C.int(a)
	bCint := C.int(b)

	// call C function sum()
	sumCint := C.sum(aCint, bCint)

	// convert back from C int to Go int
	sum := int(sumCint)
	fmt.Printf("%d + %d = %d\n", a, b, sum)
}
