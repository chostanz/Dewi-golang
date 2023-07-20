package main

import (
	"fmt"
	"strconv" //buat convert int ke string
)

func multiply(angka1 int, angka2 int) int{
	return angka1 * angka2
}

func getBiography(umur int, nama string, status string)string{
	age := strconv.Itoa(umur) //membuat var lalu di convert
	return nama + " adalah seorang " +status+ " saat ini berumur "+age+ " tahun "
}

//dua nilai string
func biography(age int, name string, statuc string)(string, string){
	ageNow := strconv.Itoa(age) //membuat var lalu di convert
	return name + " adalah seorang " +statuc,
	 " saat ini berumur "+ageNow
}

func main(){
	fmt.Println("fungsi ini menghasilkan", multiply(1,4))

	fmt.Println(getBiography(18, "arul","hehe"))


	basicInfo, ageInfo := biography(20, "sanji", "pemain basket")
	fmt.Println(basicInfo, ageInfo)
}
