package main

import "fmt"

func factorialRec(value int)int{
	//KONDISI BERHENTINYA
	if value == 1{
		return 1
	}else {
		//seakan2 5 *4, *3 *2 *!
		return value * factorialRec((value-1))
	}
}
func main() {
	fmt.Println(factorialRec(5))
}
