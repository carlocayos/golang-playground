package main

import (
	"fmt"
)

func main() {

	fmt.Println("Loops and Conditions")

	// 0 1 2 3 4 5 6 7 8 9
	for i:=0; i<10; i++ {
		fmt.Printf("%d ", i)
	}

	// using range
	slice1 := []string{"dog", "cat", "bird", "possum"}
	for i, animal := range slice1 {
		fmt.Println(i, animal)
		// 1 cat
		// 2 bird
		// 3 possum
	}


}
