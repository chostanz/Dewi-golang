package main

import "fmt"

func bayarHutank(uang int) (pesan string){

	utang :=90000
	switch {
	case uang > utang:
		pesan = "woy"
	case uang == utang:
		pesan ="mada kono sekaiwa"
	case uang < utang:
		pesan = "anjay wibu"
	}
	return
}

func cekLampu(warna string)(pesan string){
	switch warna{
		case "merah":
			pesan ="Berhenti"
		case "kuning":
			pesan ="Pelan-pelan pak supir"
		case "hijau":
			pesan="trobos ajalah anyinglah"
		default:
			pesan="warna apa ini???"

	}
	return
}

//untuk mengubah suatu fungsi walau scopenya ga di global, jd bs mengubah nilai aslinya
func pointer(point *int){
	*point =100

}

func main(){
	fmt.Println(bayarHutank(50000))
	fmt.Println(cekLampu("hijau"))

	var point = 200

	pointer(&point)

	fmt.Print(point)

}