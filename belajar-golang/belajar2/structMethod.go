package main

import "fmt"

type Pelanggan struct {
	Nama, Alamat string
	Umur         int
}

//functionnya namanya di tengah tuh
func (customer Pelanggan) sayHi(nama string) {
	fmt.Println("Hi", nama, "my name is", customer.Nama)
}

func (a Pelanggan) sayHem() {
	fmt.Println("PP", a.Nama)
}

func main() {
	dewo := Pelanggan{"Dewier", "bendungab", 29}
	// sayHi(dewo, "HEHE")
	dewo.sayHi("Ikham")
	dewo.sayHem()
}
