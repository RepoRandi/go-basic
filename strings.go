package main

import (
	"fmt"
	"strings"
)

func main50() {
	fmt.Println(strings.Contains("Randi Maulana", "Maulana"))
	fmt.Println(strings.Split("Randi Maulana Akbar", " "))
	fmt.Println(strings.ToLower("Randi Maulana Akbar"))
	fmt.Println(strings.ToUpper("Randi Maulana Akbar"))
	fmt.Println(strings.Trim("Randi Maulana Akbar", ""))
	fmt.Println(strings.TrimSpace("Randi Maulana Akbar"))
	fmt.Println(strings.TrimPrefix("Randi Maulana Akbar", "Randi"))
	fmt.Println(strings.TrimSuffix("Randi Maulana Akbar", "Akbar"))
	fmt.Println(strings.Replace("Randi Maulana Akbar", "Maulana", "Maulana Akbar", -1))
	fmt.Println(strings.ReplaceAll("Randi", "Maulana", "Maulana Akbar"))
	fmt.Println(strings.Repeat("Randi Maulana Akbar", 3))
	fmt.Println(strings.Count("Randi Maulana Akbar", "Maulana"))
	fmt.Println(strings.Index("Randi Maulana Akbar", "Maulana"))
	fmt.Println(strings.LastIndex("Randi Maulana Akbar", "Maulana"))
	fmt.Println(strings.IndexAny("Randi Maulana Akbar", "Maulana"))
	fmt.Println(strings.LastIndexAny("Randi Maulana Akbar", "Maulana"))
	fmt.Println(strings.HasPrefix("Randi Maulana Akbar", "Randi"))
	fmt.Println(strings.HasSuffix("Randi Maulana Akbar", "Akbar"))
	fmt.Println(strings.Join([]string{"Randi", "Maulana", "Akbar"}, " "))
	fmt.Println(strings.NewReader("Randi Maulana Akbar"))
	fmt.Println(strings.ToTitle("Randi Maulana Akbar"))
	fmt.Println(strings.TrimSpace(" \t\nRandi Maulana Akbar \t\n"))
	fmt.Println(strings.TrimLeft(" \t\nRandi Maulana Akbar \t\n", " \t\n"))
	fmt.Println(strings.TrimRight(" \t\nRandi Maulana Akbar \t\n", " \t\n"))
	fmt.Println(strings.Trim(" \t\nRandi Maulana Akbar \t\n", " \t\n"))
	fmt.Println(strings.TrimFunc(" \t\nRandi Maulana Akbar \t\n", func(r rune) bool {
		return r == ' ' || r == '\t' || r == '\n'
	}))
	fmt.Println(strings.Fields(" \t\nRandi Maulana Akbar \t\n"))
	fmt.Println(strings.FieldsFunc(" \t\nRandi Maulana Akbar \t\n", func(r rune) bool {
		return r == ' ' || r == '\t' || r == '\n'
	}))
	fmt.Println(strings.IndexRune(" \t\nRandi Maulana Akbar \t\n", ' '))
}
