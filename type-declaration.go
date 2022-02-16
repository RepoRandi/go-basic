package main

import "fmt"

func main() {
	type Name string
	type Married bool

	var name Name = "Randi"

	var marriedStatus Married = false

	fmt.Println(name)
	fmt.Println(marriedStatus)
}
