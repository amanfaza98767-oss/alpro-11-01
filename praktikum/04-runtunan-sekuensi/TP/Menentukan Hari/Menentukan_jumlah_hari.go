package main
import ("fmt")

func main() {
	var tahun int
	var tahunKabisat bool
	var bulan string
	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&tahun)
	fmt.Print("Masukkan bulan: ")
	fmt.Scan(&bulan)

	switch bulan {
	case "Jan", "Mar", "Mei", "Jul", "Agu", "Okt", "Des":
		fmt.Println("Jumlah hari: 31")
	case "Apr", "Jun", "Sep", "Nov":
		fmt.Println("Jumlah hari: 30")
	case "Feb":
		if (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0) {
			tahunKabisat = true
		} else {
			tahunKabisat = false
		}
		if tahunKabisat {
			fmt.Println("Jumlah hari: 29")
		} else {
			fmt.Println("Jumlah hari: 28")
		}
}
}