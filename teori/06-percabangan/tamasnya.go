package main
import "fmt"
func main() {
	var N, mobil uint8
	fmt.Scan(&N)

	mobil = N / 7 // 10 / 7 -> 3

	if N%7 !=0 {
		mobil++
	}
	fmt.Println("jumlah mobil yang diperlukan adalah", mobil)
}