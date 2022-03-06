package main

import (
	"fmt"
	"math"
)

func main52() {
	round := math.Round(1.6)
	fmt.Println(round)

	floor := math.Floor(1.6)
	fmt.Println(floor)

	ceil := math.Ceil(1.4)
	fmt.Println(ceil)

	max := math.Max(1.6, 2.4)
	fmt.Println(max)

	min := math.Min(1.6, 2.4)
	fmt.Println(min)
}
