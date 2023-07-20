package main

import "fmt"

//closure adlh kemampuan function berinteaksi dengan data2 disekitarnya dalam scope yg sa,a
//scope tu lingkupya
//closure hrus digunakan dgn bijak
//bisa ngakses data disekitarnya yaitu di main

func main() {
	counter := 0
	name := "moli"
	//membuat function increment
	increment := func() {
		//nah ini contoh merubah value dr variabel. gaboleh
		name = "budi"

		name1 := "hehe"
		fmt.Println("Penambahan")
		//hati2 membuat function di dlm function jgn sampe merubah counter diatas
		counter++
		fmt.Println(name1)
	}

	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
