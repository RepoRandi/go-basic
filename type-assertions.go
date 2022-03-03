package main

func Random() interface{} {
	return true
}

func main41() {

	result := Random()

	if _, ok := result.(string); ok {
		println("It's a string")
	} else if _, ok := result.(int); ok {
		println("It's an int")
	} else if _, ok := result.(bool); ok {
		println("It's a bool")
	} else {
		println("It's something else")
	}
}
