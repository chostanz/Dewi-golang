package main

import "fmt"

func main(){
	//break menghentikan perulangan
	//continue menghentikan perulangan dan meneruskan ke perulangan selnajutnya

	for i := 0; i <10; i++{
		if i ==5 {
			break
		}
		fmt.Println("perulangan ke", i)
	}
	for i := 0; i <10; i++{
		 if i % 2 ==0{
			continue
		}
		fmt.Println("perulangan ke", i)
	}
}