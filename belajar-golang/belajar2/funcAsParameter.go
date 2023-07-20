package main

import "fmt"

//argumennya bertipe function namanya filter parameternya string dan return valuenya string
func sayHello(name string, filter func(string) string ){
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spanFilter(name string)string{
	if name == "anjing"{
		return "dog"
	}else{
		return name
	}
}
func main() {
	//function bisa digunakan sebagai parameter untuk function lain
	sayHello("dewi", spanFilter)
	sayHello("anjing", spanFilter)
	



}
