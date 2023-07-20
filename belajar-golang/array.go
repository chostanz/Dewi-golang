package main
import "fmt"

func main(){
	//array berisi data dgn tipe data yg sama
	//daya tampung array tidak bisa bertambah setelah array dibuat

	var nama = [3]string{"tragalgar", "D", "Wotar"}


	var wibu [3] string
	wibu[0] = "HAHA"
	wibu[1] ="FUFUF"
	wibu[2]="xixixi"

	//mengganti nilai di salah satu
	nama[1] = "w"
	fmt.Println(nama)

	fmt.Println(wibu[2])


	//function len(array) untuk menghitung panjang array
	fmt.Println(len(nama))
	fmt.Println(len(wibu[2]))
	
}