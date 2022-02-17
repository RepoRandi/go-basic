package main

import "fmt"

func main12() {
	var names [3]string

	names[0] = "John"
	names[1] = "Paul"
	names[2] = "George"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		10,
		20,
		30,
	}

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println("length:", len(values))

	consoleGames := [...]string{
		"Super Mario Bros.",
		"Tetris",
		"The Legend of Zelda",
		"Batman",
	}

	for _, game := range consoleGames {
		fmt.Println(game)
	}

	fmt.Println("length:", len(consoleGames))
}
