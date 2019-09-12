package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Channel")

	c := make(chan int)

	go sum(c, 5, 7)
	go sum(c, 2, 3)

	x, y := <-c, <-c
	fmt.Printf("x = %d, y = %d", x, y)

	// x = 5, y = 12 - unordered
	// might be x = 12, y = 5

}

var count int = 1

func sum(c chan int, x, y int) {

	if count == 1 {
		time.Sleep(1 * time.Second)
		count += 1
	}

	c <- x + y
}
