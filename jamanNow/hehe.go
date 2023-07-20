package main

import "fmt"

func main() {
	fmt.Println("Benar ", true)

	//menghitung jumlah string pakai len
	fmt.Println(len("Dewik"))
	fmt.Println("Dewik"[0])

	//variabel
	var age int8 = 34
	var umur int = 9
	fmt.Println(age, umur)

	var (
		firstName = "Dewi "
		lastName  = "Rahmawati"
	)

	fmt.Println(firstName, lastName)

	//konstantan harus lgsg dikasih datanya.
	const nama string = "wibu"
	const (
		name string = "hehe"
		hehe        = "wooo"
	)
	fmt.Println(nama, name, hehe)

}
