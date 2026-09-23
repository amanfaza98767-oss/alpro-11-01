package main
import "fmt"

func main() {
	var nilai1, nilai2 int

	//Membaca input dari user
	fmt.Scan(&nilai1)
	fmt.Scan(&nilai2)

	//menukar nilai
	nilai1, nilai2 = nilai2, nilai1

	//Menampilkan hasil
	fmt.Println("Setelah ditukar:")
	fmt.Println("Nilai 1:", nilai1)
	fmt.Println("Nilai 2:", nilai2)
}