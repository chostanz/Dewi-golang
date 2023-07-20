package main

import "fmt"

func main() {
	//operasi boolean, && and, || or, ! kebalikannya
	//operasi && jika true true = true, selainnya false
	//operasi || jika false false = false, selainnya true
	//operasi ! tu jadi jika true = false, false = true
	var ujian = 80
	var absensi = 80

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80

	var lulus = lulusUjian && lulusAbsensi

	fmt.Println(lulus)

	//yg bagus nulisnya kek gini
	fmt.Println(ujian > 80 || absensi > 80)
}
