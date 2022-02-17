package main

func main() {
	switch name := "Lad"; name {
	case "Randi":
		println("Hello", name)
	case "Lady":
		println("Hello", name)
	default:
		println("Hello stranger")
	}
}
