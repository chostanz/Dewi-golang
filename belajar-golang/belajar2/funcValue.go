package main
import "fmt"


func sayGoodBye(name string)string{
	return "Good bye"+name
}
func main(){
	//function merupakan tipe data dan bisa disimpan dalam variabel

	//membuat variabel yg isinya function
	goodbye := sayGoodBye

	result := goodbye("Dewi")

	fmt.Println(result)
	fmt.Println(sayGoodBye("Sanji"))

}