package main

import "fmt" 

func main() {
	//slice adalah potongan dari array, mirip array. ukuran slice bisa berubah
	//tipe data slice ada 3 data, pointer, length, capacity
	//pointer adlh petunjuk data pertama di array para slice
	//length adalah panjang dari slice
	//capacity adalah kapasitas dari slice, length tidak boleh lebih dari capacity

	//array[low:high] ->membuat slice dr low ke sblm high
	//array[low:] -> membuat slice dimulai dr low sampai akhir
	//array[:high] -> membuat slice dimulai index 0 smp index sblm high
	//array[:] -> membua slice dimulai index 0 sampai index akhir

	var months =[...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	//cap tu capacity dari 7 sampe 12
	fmt.Println(cap(slice1))

	slice1[0] = "Mei lagi"
	fmt.Println(months)

	//append menambah data ke posisi terakhir data
	//append (slice, data)

	//make([]tipedata, length, capacity) membuat slice baru

	//copy(destination,source) menyalin source ke destination


	days := [...]string{"Senin", "selasa", "rabu", "kamis", "jumat", "sabtu","minguu"}

	daySlice1 := days[2:]

	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"

	fmt.Println(daySlice1) //sabtu baru minggu baru

	daySlice2 := append(daySlice1, "Libur baru")  //akan membuat sebuah array baru krn sdh penuh
	daySlice2[0] = "UPS"
	fmt.Println(daySlice2)
	fmt.Println(days)



	//membuaft slice dari awal
	newSlice := make([]string, 2 , 5)
	newSlice[0] = "dewo"
	newSlice[1] = "rahmawat"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))


	//copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))

	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	//beda array dan slice
	kiArray := [...]int{1,2,3,4}
	kiSlice := []int{1,2,3,4,5}

	fmt.Println(kiArray)
	fmt.Println(kiSlice)

}
