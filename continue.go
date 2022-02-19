package main

func main() {
	for i := 0; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}
		println("Index:", i)
	}
}
