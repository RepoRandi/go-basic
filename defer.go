package main

import "fmt"

func main31() {
	runApplication(0)
}

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Sedang menjalankan aplikasi")

	result := 10 / value
	fmt.Println("Result:", result)
}
