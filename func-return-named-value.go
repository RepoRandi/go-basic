package main

func main24() {
	firstName, middleName, lastName := getFullName2()
	println(firstName, middleName, lastName)
}

func getFullName2() (firstName, middleName, lastName string) {
	return "John", "Paul", "Jones"
}
