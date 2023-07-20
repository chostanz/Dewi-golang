package main

import "fmt"

//argumennya bertipe function namanya filter parameternya string dan return valuenya string
// func sayHello(name string, filter func(string) string ){
// 	nameFiltered := filter(name)
// 	fmt.Println("Hello", nameFiltered)
// }


//karna kepanjangan nulis function di argumen, jadi kt bisa deklarasi tipenya dulu
type Filter func(string) string

//Filter tu tipe data alias dari function
func getHello(name string, filter Filter){
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string)string{
	if name == "anjing"{
		return "dog"
	}else{
		return name
	}
}
func main() {
	//function bisa digunakan sebagai parameter untuk function lain
	getHello("dewi", spamFilter)
	getHello("anjing", spamFilter)
	



}
