package main

import "fmt"

func main() {
	//mengaliaskan tipe data
	//jadi noKTP alias untuk string
	type noKTP string
	type married bool

	var ktpDewi noKTP = "340411610"
	var status married = true
	fmt.Println(ktpDewi)
	fmt.Println("status menikah : ", status)

}

