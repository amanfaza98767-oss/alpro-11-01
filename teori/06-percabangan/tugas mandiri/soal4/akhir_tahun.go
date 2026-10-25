package main
import "fmt"

func main() {
	var belanja int
	var bersedia bool
	fmt.Scan(&belanja, &bersedia)

	kartu := bersedia
	diskon := belanja >= 100000
	casback := belanja >= 200000 && kartu

	bayar := belanja

	if diskon {
		bayar -= belanja / 10 
	}
	if casback {
		bayar -= 75000
	}

	fmt.Println("Kartu?", kartu)
	fmt.Println("Diskon?", diskon)
	fmt.Println("Cashback?", casback)
	fmt.Printf("Rp. %d\n", bayar)

}