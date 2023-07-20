package main

import "fmt"

func main() {

	no := 1

	for no <=10{
	fmt.Println("Perulangan ke", no)
	no++
	}

	for angka :=1; angka <5 ; angka++{
		fmt.Println(angka)
	}


	//for range, unutk interaksi data yg ada di array
	 var slice =[]string{"heri","milo","moli"}

	 for i := 0; i< len(slice); i++{
		fmt.Println(slice[i])
	 }

	 //range
	 for index, name := range slice{
		fmt.Println("index", index, "=", name)
	 }


	 //map di for

	 dataSiswa := make(map[string]string)
	 dataSiswa["nama"] = "hendri"
	 dataSiswa["nomor"] = "19007"

	 for key, value := range dataSiswa{
		fmt.Println(key, "=", value)
	 }
}