package main

import "fmt"

func main(){
	//map bisa menenntukan index sesuka hati kita
	//kunci gaboleh sama,  unik

	//jumlah data yg bmasuk map boleh sebanyak2nya dan ga harus urut

	dataSiswa := map[string]string{
		"nama": "Dewi",
		"kelas" : "13 SIJA A",
	}

	dataSiswa["alamat"] = "Bendungan"
	
	fmt.Println(dataSiswa)
	fmt.Println(dataSiswa["nama"])
	fmt.Println(dataSiswa["kelas"])

	//menambah data
	

}