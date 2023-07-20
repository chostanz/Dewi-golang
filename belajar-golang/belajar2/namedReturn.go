package main

import "fmt"

func getName() (first, middle, last string) {
	first = "kucel "
	middle = "nisrina "
	last = "ali "

	return
}
func main() {
	a, b, c := getName()
	fmt.Print(a, b, c)
}
