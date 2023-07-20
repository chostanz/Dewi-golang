package main

import "fmt"

//recover adalah function untuk menangkap data panic
func endApl() {
	fmt.Println("Aplikasi selesai")
	//RECOVER HARUS DI SIMPAN DI DEFER PUNCTION. nek setelahnya keburu panik jd ga ngaruh
	message := recover()
	fmt.Println("Terjadi error: ", message)
}
func runAplikasi(error bool) {
	defer endApl()

	if error {
		panic("error")
	}
	fmt.Println("Aplikasi berjalan")
}
func main() {
	runAplikasi(true)
	//recovery tu menghentikan panicnya jd program dibawahnya masih bs jalan
	fmt.Println("dewi")
}
