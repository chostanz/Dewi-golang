package main

import "fmt"

//struct adalah template data yg digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
//pokoke kek array tp tipe datanya beda2
//data struct disimpan dalam field. sturct = kumpulan field

//membuat structnya depannya uppercase
type Customer struct {
	//Name string
	Name, Address string
	Age           int
}

//struct tdk bs menggunakan langsung
//namun kita bisa membuat data/object dr struct yg telah dibuat

func main() {
	var data Customer
	data.Name = "Dewi"
	data.Address = "Bendungan"
	data.Age = 18

	fmt.Println(data)

	//cara kedua bikin struct
	biodata := Customer{
		Name:    "Dewi",
		Address: "12 SIJA A",
		Age:     18,
	}
	fmt.Println(biodata)

	//cara ketiga
	//tidak efektif, rentan error
	budi := Customer{"Budi", "jogja", 90}
	fmt.Println(budi)

}
