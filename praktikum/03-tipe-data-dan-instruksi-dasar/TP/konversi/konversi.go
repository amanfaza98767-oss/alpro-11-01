package main

import "fmt"

func main() {
	var f float64
	fmt.Print("Masukkan nilai f: ")
	fmt.Scan(&f)

	const mil = 1.6
	km := f * mil
	fmt.Printf("%.1f mil = %.1f km", f, km)
}

