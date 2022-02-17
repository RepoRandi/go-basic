package main

func main15() {
	name := "Lady"

	if name == "Randi" {
		println("Hello", name)
	} else if name == "Lady" {
		println("Hello", name)
	} else {
		println("Hello stranger")
	}

	if len(name) > 5 {
		println(name, "is a long name")
	} else {
		println(name, "is a short name")
	}
}
