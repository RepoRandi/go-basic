package main

func main1() {
	switch name := "Lad"; name {
	case "Randi":
		println("Hello", name)
	case "Lady":
		println("Hello", name)
	default:
		println("Hello stranger")
	}
}
