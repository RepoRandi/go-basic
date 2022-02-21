package main

func main30() {
	counter := 0

	increment := func() {
		println("Increment")
		counter++
	}

	increment()
	increment()

	println(counter)
}
