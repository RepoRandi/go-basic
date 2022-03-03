package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main42() {
	var address1 Address = Address{City: "Toronto", Province: "Ontario", Country: "Canada"}
	var address2 *Address = &address1

	address2.City = "Ottawa"

	*address2 = Address{City: "Montreal", Province: "Quebec", Country: "Canada"}

	fmt.Println(address1)
	fmt.Println(address2)

	var address3 *Address = new(Address)
	// address3.City = "Vancouver"
	// address3.Province = "British Columbia"
	// address3.Country = "Canada"

	fmt.Println(address3)
}
