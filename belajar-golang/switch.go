package main

import "fmt"

func main() {
	var nama = "Dewi"

	switch nama {
	case "Sanji":
		fmt.Println("Halo Sanji")
	case "Dewi":
		fmt.Println("Hallo Dewi")
	case "Zoro":
		fmt.Println("Hallo zoro")
	}


	//switch short statement
	//nilainya boolean
	//var length = len(nama)
	switch length := len(nama); length >5{
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah cukup")
	}


		//switch tanpa kondisi
		length2 := len(nama)
		switch {
		case length2 > 10:
			fmt.Println("Nama terlalu panjang")
		case length2 == 5:
			fmt.Println("Nama sudah pas")
		case length2 < 5:
			fmt.Println("Nama kurang")
		}
}
