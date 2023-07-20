package main

import "fmt"

func endApp() {
	fmt.Println("app selesai")
}

func runApl(error bool) {
	defer endApp()

	if error {
		//panic digunakan untuk memberhentikan program saat error, namun defer ttp berjalan
		panic("Aplikasi error")
	}
	fmt.Println("aplikasi berjalan")
}
func main() {
	runApl(true)
}
