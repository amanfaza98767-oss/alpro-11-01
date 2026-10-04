package main

import "fmt"

func main() {
	var y, x int
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	sisa := y % x
	fmt.Println("Hasil Sisa Bagi:", sisa)
}