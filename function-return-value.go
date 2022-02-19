package main

func main22() {
	s, i := getSample("Juragan")
	println(s, i)
}

func getSample(name string) (string, int) {
	return "Hello" + " " + name, 99
}
