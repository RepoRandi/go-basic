package main

import "fmt"

func main10() {
	var name1 = "Randi"
	var name2 = "Randi"

	var result bool = name1 == name2

	fmt.Printf("%v\n", result)

	var value1 = 12
	var value2 = 10

	if value1 > value2 {
		fmt.Printf("%v is greater than %v\n", value1, value2)
	} else {
		fmt.Printf("%v is less than %v\n", value1, value2)
	}
}
