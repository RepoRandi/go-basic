package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return 2
	} else {
		return "Ups"
	}
}

func main38() {
	var data interface{} = Ups(5)

	fmt.Println(data)
}
