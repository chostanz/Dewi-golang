package main

import "fmt"

func main() {
	var name = "Dewi"

	if name == "Dewi" {
		fmt.Println("Hello", name)

	} else {
		fmt.Println("Nama tidak terdaftar")
	}

	//kode program if short statement
	//var length = len(name)
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}


}
