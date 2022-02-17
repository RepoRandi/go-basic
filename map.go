package main

import "fmt"

func main15() {
	person := map[string]string{
		"name": "John",
		"job":  "Software Engineer",
	}

	person["email"] = "randi@gmail.com"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["job"])
	fmt.Println(person["email"])

	var book map[string]string = make(map[string]string)

	book["name"] = "The Go Programming Language"
	book["author"] = "Alan A. A. Donovan"
	book["publisher"] = "Addison-Wesley Professional"

	fmt.Println(book)

	delete(book, "publisher")

	fmt.Println(book)
}
