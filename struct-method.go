package main

import "fmt"

type Customers struct {
	Name, Address string
	Age           int
}

func main36() {
	var customer Customers
	customer.Name = "John"
	customer.Address = "123 Main St"
	customer.Age = 20
	customer.sayHello()
}

func (customers Customers) sayHello() {
	fmt.Println("Hello", customers.Name)
}
