package main

import "fmt"

//interface isinya definisi2 method

//membuat kontrak baru
type HasName interface {
	GetName() string
}

func SayHelo(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Namaa string
}

func (person Person) GetName() string {
	return Namaa
}

//implementasi interface scr otomatis

func main() {
	SayHelo()
}
