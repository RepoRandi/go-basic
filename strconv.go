package main

import (
	"fmt"
	"strconv"
)

func main51() {
	boolean, err := strconv.ParseBool("false")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(boolean)
	}

	number, err := strconv.ParseInt("123", 10, 64)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(number)
	}

	float, err := strconv.ParseFloat("123.456", 64)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(float)
	}

	string := strconv.FormatBool(boolean)
	fmt.Println(string)

	string = strconv.FormatInt(number, 10)
	fmt.Println(string)

	string = strconv.FormatFloat(float, 'f', 2, 64)
	fmt.Println(string)

	string = strconv.Itoa(123)
	fmt.Println(string)

	valueInt, err := strconv.Atoi("123")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valueInt)
	}

	valueFloat, err := strconv.ParseFloat("123.456", 64)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valueFloat)
	}

}
