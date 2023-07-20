package main

import "fmt"

var text = "ini adalah teks rahasia"

//fungsi tanpa return
func getText(){
	fmt.Print(text)
}

func main(){
	getText();
}