package main

import "fmt"

//func yg tida memiliki nama function
 type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist){
	if blacklist(name){
		fmt.Println("Kamu diblokir", name)
	}else{
		fmt.Println("Welcome", name)
	}
}

// func blacklist(name string)bool{
// 	return name := "hehe"
// }


func main(){
	blacklist := func(name string) bool{
		return name =="admin"
	}

	registerUser("admin",blacklist)

	registerUser("root", blacklist)
}