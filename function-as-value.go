package main

func main26() {
	goodbye := getGoodbye

	println(goodbye("World"))
}

func getGoodbye(name string) string {
	return "Goodbye" + " " + name
}
