package main

func main32() {
	runApp(true)
}

func endApp() {
	println("Aplikasi Selesai")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Terjadi kesalahan")
	}

	println("Aplikasi Sedang Berjalan")
}
