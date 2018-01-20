//Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.
//You might find strings.Fields helpful.

package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {

	stringArray := strings.Fields(s)

	fmt.Println(stringArray)
	fmt.Println(len(stringArray))

	m := make(map[string]int)

	for _, val := range stringArray {
		count := m[val]
		m[val] = count + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
