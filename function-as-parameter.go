package main

import "strings"

type Filter func(string) string

func main27() {
	sayHelloWithFilter("Anjing", spamFilter)
	sayHelloWithFilter("Jancok", spamFilter)
	sayHelloWithFilter("Bambang", spamFilter)
	sayHelloWithFilter("big e", uppercaseFilter)
}

func sayHelloWithFilter(name string, filter Filter) {
	println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "A....g"
	} else if name == "Jancok" {
		return "J....k"
	} else {
		return name
	}
}

func uppercaseFilter(name string) string {
	return strings.ToUpper(name)
}
