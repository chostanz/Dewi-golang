package main

import "fmt"

//konfersi tipe data dari int ke int
func main() {
	var nilai32 int32 = 6374
	var nilai64 int64 = int64(nilai32)

	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)



	//konversi tipe data byte ke string

var name = "Dewi R"
var e = name[0]  //uint 8 alias byte

var eString string = string(e)
fmt.Println(name)
fmt.Println(eString)




}
