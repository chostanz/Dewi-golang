package main

import "fmt"

func hello() {
	fmt.Println("hello")

	for i:=0; i>10;i++{
		fmt.Println(i)
	}
}

func main() {
	hello()
}
