package main
import "fmt"

func main(){
	//operasi  perbandingan
	var name1 = "Dewi"
	var name2 = "Dewi2"

	var result bool = name1==name2

	fmt.Println(result)

	var value1 = 100
	var value2 = 200
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}