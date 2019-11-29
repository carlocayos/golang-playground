package main

import (
	"flag"
	"fmt"
)

func main() {
	result := flag.String("endpoint", "This is the Default Value", "Example string flag")
	flag.Parse()
	fmt.Printf("%s\n", *result)
}
