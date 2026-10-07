package main
import (
"fmt"
)

func main() {
	var pajak float64
	var gaji float64

	fmt.Scan(&gaji)
	
if gaji <= 50 {
		pajak = 0.05 * gaji
	} else if gaji > 50 && gaji <= 100 {
		pajak = 0.05 * 50
		pajak += 0.1 * (gaji - 50)
	} else if gaji > 100 && gaji <= 200 {
		pajak = 0.05 * 50
		pajak += 0.1 * 50
		pajak += 0.15 * (gaji - 100)
	} else if gaji > 200 {
		pajak = 0.05 * 50
		pajak += 0.1 * 50
		pajak += 0.15 * 100
		pajak += 0.2 * (gaji - 200)
	}
	fmt.Printf("Pajak yang harus dibayarkan: %.2f", pajak)

}