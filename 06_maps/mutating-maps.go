package main

import (
	"fmt"
)

const NameKey = "name"

func main() {
	m := make(map[string]string)

	m[NameKey] = "Carlo"

	// get the value from key
	fmt.Println(m[NameKey])

	// update the value
	m[NameKey] = "CarloUpdated"
	fmt.Println(m[NameKey])

	// check if key exists
	value, ok := m[NameKey]
	fmt.Printf("value=%s, exists=%t \n", value, ok)
	fmt.Println(ok)

	// delete - returns no value
	delete(m, NameKey)
	fmt.Println(m[NameKey])


}
