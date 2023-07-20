package main

import "fmt"

//defer tu function yg bisa dijadwalkan untuk diesekusi setelah sebuuah function selesai dieksekusi

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApp(value int) {
	//defer harus dideklarasikan diatas
	defer logging()
	result := 10 % value
	fmt.Println("Running app", result)
}

func main() {
	runApp(3)
}
