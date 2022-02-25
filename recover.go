package main

import "fmt"

func main33() {
	runApps(true)
}

func endApps() {
	message := recover()

	if message != nil {
		fmt.Println("Error:", message)
	}

	fmt.Println("Aplikasi Selesai")
}

func runApps(err bool) {
	defer endApps()

	if err {
		panic("Terjadi kesalahan")
	}

	fmt.Println("Aplikasi berjalan")
}
