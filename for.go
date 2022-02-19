package main

func main17() {

	for counter := 1; counter <= 10; counter++ {
		println(counter)

	}

	consoleGames := []string{
		"Super Mario Bros.",
		"Contra",
		"Metroid",
		"Mega",
		"Pac",
		"Tetris",
		"Zelda",
	}

	for _, game := range consoleGames {
		println("Game: " + game)
	}

	persons := make(map[string]string)
	persons["first"] = "John"
	persons["last"] = "Doe"
	persons["age"] = "25"

	for key, person := range persons {
		println(key, ":", person)
	}
}
