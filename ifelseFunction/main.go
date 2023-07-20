package main

import "fmt"


func bayarUtang(uang int)(pesan string){
	//utang := 50000

	if utang :=50000; uang> utang{
		pesan = "Uangmu kebanyakan"
	}else if uang == utang {
		pesan = "uangmu pas"
	}else {
		pesan = "kurang anjay"
	}

	return
}

func main(){
	fmt.Println(bayarUtang(1000))
}