package main
import "fmt"
func main() {
	var t int
	fmt.Print("input t (jumlah transaaksi): ")
	fmt.Scan(&t)

	totalpendapatan := 0
	for i := 1; i <= t; i++ {
		var jenisfilm, hari string
		var jumlahtiket int
		fmt.Print("Input transaksi ke- %d (jenis, hari, tiket): ", i)
		fmt.Scan(&jenisfilm, &hari, &jumlahtiket)

		hargaSatuan := 0

	
		if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" ||  {
			if jenisfilm == "Reguler" {
				hargaSatuan = 40000
			} else if jenisfilm == "3D" {
				hargaSatuan = 50000
			} else if jenisfilm == "IMAX" {
				hargaSatuan = 80000
			}
		}

		else if hari == "Jumat" || hari == "Sabtu" || hari == "Minggu" {
			if jenisfilm == "Reguler" {
				hargaSatuan = 50000
			} else if jenisfilm == "3D" {
				hargaSatuan = 60000
			} else if jenisfilm == "IMAX" {
				hargaSatuan = 90000
			}
		}
		totalHarga := hargaSatuan * jumlahTiket
		
		if jumlahTiket >= 5 {
			diskon := (totalHarga * 10) / 100
			totalHarga = totalHarga - diskon
		}

		fmt.Printf ("output total harga transaksi ke-%d: %d\n", i, totalHarga)
		totalPendapatan = totalPendapatan + totalHarga
	}
	fmt.Println("output total pendapatan:", totalpendapatan)
}