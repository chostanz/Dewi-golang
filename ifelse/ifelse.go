package main
import "fmt"


func main(){
	num := 10

	if num>15{
		fmt.Println("Kerja bagus")
	}else{
		fmt.Println("hehooo")
	}



	color :="hijaau"

	if color == "hijau"{
		fmt.Println("Meletus balon hijau")
	}else{
		fmt.Println("Terkejut")
	}

	utang:= 10000
	uang := 10000

	if uang > utang{
		fmt.Print("Bayar hutang")
	}else if uang == utang{
		fmt.Print("uang pas")
	}

}