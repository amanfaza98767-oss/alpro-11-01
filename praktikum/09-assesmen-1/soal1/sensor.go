package main
import "fmt"

func main() {
	var sekarang, sebelumnya, tertinggi, terendah, total, jumlahData int
	var rataRata float64
	
	fmt.Scan (&sebelumnya)

	tertinggi = sebelumnya
	terendah = sebelumnya
	total = sebelumnya
	jumlahData = 1

	for {
		fmt.Scan(&sekarang)

		if sekarang == sebelumnya{
		break
		}
		if sekarang > tertinggi {
			tertinggi = sekarang
		}
		if sekarang < sebelumnya {
			sebelumnya = sekarang
		}
		total += sekarang
		jumlahData++

		sebelumnya = sekarang
	}
	rataRata = float64(total) / float64(jumlahData)
	fmt.Printf("%d %d %.2f\n", tertinggi, terendah, rataRata)
}