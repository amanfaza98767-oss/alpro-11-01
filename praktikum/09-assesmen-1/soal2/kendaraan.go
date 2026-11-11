package main
import "fmt"

func main() {
	var kendaraan string
	var motor, mobil, truk int

	for {
		fmt.Scan(&kendaraan)

		switch kendaraan {
		case "M":
			motor++
		case "K":
			mobil++
		case "T":
			truk++
		default:
		break
		}
		if kendaraan != "M" && kendaraan != "K" && kendaraan != "T"{
			break
		}
	}
	fmt.Println("Tipe M =", motor)
	fmt.Println("Tipe K =", mobil)
	fmt.Println("Tipe T =", truk)
	
}