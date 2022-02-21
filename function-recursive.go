package main

func main() {
	println(factorialLoop(5))
	println(factorialRecursive(5))
}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 0 {
		return 1
	}
	return value * factorialRecursive(value-1)
}
