package main

import "fmt"

func main() {
	var angka int
	fmt.Scan(&angka)

	if angka == 0 {
		fmt.Print("Nol")
	} else {
		if angka%2 == 0 {
			fmt.Print("Genap, ")
		} else {
			fmt.Print("Ganjil, ")
		}

		if angka > 0 {
			fmt.Println("Positif")
		} else {
			fmt.Println("Negatif")
		}
	}
}
