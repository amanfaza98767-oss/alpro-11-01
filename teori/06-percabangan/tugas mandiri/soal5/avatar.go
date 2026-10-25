package main

import (
	"fmt"
)

func main() {
	var p int
	fmt.Scan(&p)

	dewasa := (p + 4) / 5
	if dewasa > 3 {
		dewasa = 3
	}

	sisa := p - dewasa *5
	if sisa < 0 {
		sisa = 0
	}
	
	kecil := 0
	if sisa > 0 {
		kecil = (sisa + 1) / 2
	}
	
	if kecil > 5 {
		kecil = 5
	}
	terangkut := dewasa * 5 + kecil * 2
	tidakBerangkat := p - terangkut

	fmt.Print("Dewasa ", dewasa)

	if kecil > 0 {
		fmt.Print(", kecil ", kecil)
	}

	if tidakBerangkat > 0 {
		fmt.Print(", dan ", tidakBerangkat, " tak berangkat")
	}
	fmt.Println()
}