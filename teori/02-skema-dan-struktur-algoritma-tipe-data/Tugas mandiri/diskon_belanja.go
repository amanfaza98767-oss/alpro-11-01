package main
import "fmt"
// kamus: harga, persen: integer, potongan, HargaAkhir: real
// algoritma: input harga, persen,
// potongan ← harga * persen / 100
// Harga_Akhir ← harga - potongan
// output Harga_Akhir
func main() {
	var harga, persen int
	var potongan, HargaAkhir float64

	fmt.Scan(&harga, &persen)

	potongan = float64(harga) * float64(persen) / 100
	HargaAkhir = float64(harga) - potongan

	fmt.Println(HargaAkhir)
}
