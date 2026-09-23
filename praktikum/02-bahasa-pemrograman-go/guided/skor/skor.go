package main
import "fmt"

func main() {
	var nama string
	var skorMatematika, skoringgris int

	//Membaca input dari user
	fmt.Scan (&nama)
	fmt.Scan (&skorMatematika)
	fmt.Scan (&skoringgris)

	//menghitung total skor & rata-rata skor
	totalSkor := skorMatematika + skoringgris
	rataRataSkor := totalSkor / 2

	//menampilkan output
	fmt.Println("Nama:", nama)
	fmt.Println("Total Skor:", totalSkor)
	fmt.Println("Rata-rata Skor:", rataRataSkor)
}