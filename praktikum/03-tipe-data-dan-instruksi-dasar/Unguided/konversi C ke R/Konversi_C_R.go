package main
import "fmt"

func main() {
	var celcius float64
	fmt.Print("Masukkan suhu dalam Celcius: ")
	fmt.Scan(&celcius)
	var reamur float64 = celcius * 4 / 5
	fmt.Println("Suhu dalam Reamur:", reamur)
}