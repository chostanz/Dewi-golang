package main

import "fmt"

//return satu data
func getHello(name string) string {
	if name == "" {
		return "Hello bro"
	} else {
		return "hello " + name
	}
	//return "Hello " + name

}

//function multiple values
func multiple() (string, int, string) {
	return "Dewi", 2, "Rahmawati"
}

func main() {
	//diantara kurung tutup sm kurawal tu tipe data returnnya
	result := getHello("moli")
	fmt.Println(result)

	fmt.Println(getHello(""))

	//panggil multiple
	//mengabaikan nilai yg ke 3 biar ga error _
	firstName, number, _ := multiple()
	fmt.Println(firstName, number)
}
