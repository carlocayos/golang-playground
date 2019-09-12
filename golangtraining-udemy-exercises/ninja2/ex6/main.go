package main

import "fmt"

const (
	year1 = 2019 + iota
	year2
	year3
	year4
)


func main() {
	fmt.Println(year1, year2, year3, year4)
}
