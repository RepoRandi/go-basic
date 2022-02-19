package main

func main23() {
	firstName, lastName, _ := getFullName()
	println(firstName, lastName)
}

func getFullName() (string, string, int) {
	return "Sukidi", "Sukirman", 99
}
