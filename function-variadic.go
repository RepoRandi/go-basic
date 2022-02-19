package main

func main25() {
	result := sumAll(1, 2, 3, 4, 5)
	println(result)

	slice := []int{1, 2, 3, 4, 5, 6}

	result = sumAll(slice...)
	println(result)
}

func sumAll(numbers ...int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum
}
