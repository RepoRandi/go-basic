package main

import (
	"errors"
	"fmt"
)

func Pembagi(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("error: Nilai b tidak boleh 0")
	}
	return a / b, nil
}

func main40() {
	var exampleError = errors.New("error message")

	if exampleError != nil {
		fmt.Println(exampleError.Error())
	} else {
		fmt.Println("No error")
	}

	var a, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	result, err := Pembagi(a, b)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Result:", a, "/", b, "=", result)
	}
}
