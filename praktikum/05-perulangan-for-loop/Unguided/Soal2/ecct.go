
package main

import "fmt"

func main() {
	var m int
	fmt.Scan(&m)

	jumlahSertifikat := 0

	for i := 1; i <= m; i++ {
		var p, v, ga, fs, ci, tc int
		fmt.Scan(&p, &v, &ga, &fs, &ci, &tc)

		total := p + v + ga + fs + ci + tc

		// Membandingkan tanpa desimal agar lebih akurat:
		// rata-rata >= 3.33
		if total*100 >= 333*6 {
			fmt.Println("Mendapat Sertifikat")
			jumlahSertifikat++
		} else {
			fmt.Println("Tidak Mendapat Sertifikat")
		}
	}

	fmt.Println(jumlahSertifikat)
}
