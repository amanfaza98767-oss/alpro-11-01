package main

import "fmt"

func main() {
	var  celcius float64
	fmt.Print("Masukkan suhu dalam Celcius: ")
	fmt.Scan(&celcius)
	var kelvin float64 = celcius + 273
	fmt.Println("Suhu dalam Kelvin:", kelvin)
}
