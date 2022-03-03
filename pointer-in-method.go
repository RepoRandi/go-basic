package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "MR. " + man.Name
}

func main44() {

	name := Man{Name: "John"}
	name.Married()

	fmt.Println(name.Name)
}
