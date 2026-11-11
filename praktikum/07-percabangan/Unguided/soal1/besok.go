package main
import "fmt"

func main() {
	var tanggal, bulan, tahun int
	fmt.Scan(&tanggal, &bulan, &tahun)

	kabisat := tahun%400 == 0 || (tahun%4 == 0 && tahun%100 !=0)

	jumlahHari := 0

	if bulan >= 1 && bulan <= 12 {
		switch bulan {
		case 1, 3, 5, 7, 8, 10, 12:
			jumlahHari = 31
		case 4, 6, 9, 11:
			jumlahHari = 30
		case 2:
			if kabisat {
				jumlahHari = 29
			} else {
				jumlahHari = 28
			}
		}
	}
	if tanggal < 1 || tanggal > jumlahHari {
		fmt.Println("Tanggal tidak valid")
	} else {
		tanggal++
		if tanggal > jumlahHari {
			tanggal = 1
			bulan++
			
			if bulan > 12 {
				bulan = 1
				
				tahun ++
			}
		}
		fmt.Printf("Tanggal valid\n%d/%d/%d\n",
		tanggal, bulan, tahun)
	}
}