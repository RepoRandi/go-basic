package main

type Addres struct {
	City, Province, Country string
}

func ChangeAddres(addres *Addres) {
	addres.City = "Toronto"
}

func main43() {
	addres := Addres{"London", "Ontario", "Canada"}
	ChangeAddres(&addres)
	println(addres.City)
}
