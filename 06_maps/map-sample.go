package main

import (
	"fmt"
)

type GPS struct {
	Lat, Long float64
}

var sampleMap map[string]GPS

func main() {
	fmt.Println("Map with struct")

	sampleMap = make(map[string]GPS)

	sampleMap["myhouse"] = GPS {
		5.5,
		8.9,
	}

	// {5.5 8.9}
	fmt.Println(sampleMap["myhouse"])

	// Invalid key {0 0} - default value
	fmt.Println(sampleMap["invalidKey"])
}
