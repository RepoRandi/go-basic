package main

import "fmt"

func main18() {

	for i := 0; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Index: ", i)
	}
}
