// http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#blocked_goroutines
// Blocked Goroutines and Resource Leaks

package main

import (
	"fmt"
	"time"
)

func main() {
	First()

	time.Sleep(60 * time.Second)
}

func First() {

	count := 20
	ch := make(chan int)

	sender := func(i int) {
		ch <- i
		fmt.Println(i)
	}

	for i := 0; i < count; i++ {
		go sender(i)
	}

	fmt.Println(<-ch)
}
