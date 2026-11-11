package main
import "fmt"

func main() {
	var kelas, hari string
	var usia int

	fmt.Scan(&kelas, &usia, &hari)
	
	var harga float64

	switch kelas {
	case "Ekonomi":
		harga = 100000
	case "Bisnis":
		harga = 200000
	case "Eksekutif":
		harga = 350000
	default:
fmt.Println("kelas tidak valid")
return
	}

	if hari == "Sabtu" || hari == "Minggu" {
		harga = harga * 1.2
	}

	if usia < 3 {
		harga = 0
	} else if usia <=12 {
		harga = harga * 0.5
	} else if usia >= 60 {
		harga = harga * 0.7
	} 
	fmt.Printf("%.0f\n", harga)
}