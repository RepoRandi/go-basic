package main

import "fmt"

func main16() {
	name := "Randi"

	switch name {
	case "Randi":
		println("Hello", name)
	case "Lady":
		println("Hello", name)
	default:
		println("Hello stranger")
	}

	switch len(name) >= 5 {
	case true:
		fmt.Println("Name is longer than 5 characters")
	case false:
		fmt.Println("Name is less than 5 characters")
	}

	length := len(name)

	switch {
	case length > 5:
		fmt.Println("Name is longer than 5 characters")
	case length < 5:
		fmt.Println("Name is less than 5 characters")
	case length == 5:
		fmt.Println("Name is 5 characters")
	}
}
