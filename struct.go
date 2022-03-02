package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main35() {
	var randi Customer

	randi.Name = "Randy"
	randi.Address = "Colombo"
	randi.Age = 23

	fmt.Println(randi)

	edy := Customer{
		Name:    "Edy",
		Age:     23,
		Address: "New York",
	}

	fmt.Println(edy)

	lady := Customer{"Lady", "Jakarta", 23}

	fmt.Println(lady)
}
