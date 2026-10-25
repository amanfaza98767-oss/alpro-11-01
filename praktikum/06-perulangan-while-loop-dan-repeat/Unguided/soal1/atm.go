package main
import "fmt"

func main() {
	var password string
	var kesempatan int = 0
	var saldo int = 10000000
	var pilihan int

	// proses login
	for password != "2007" && kesempatan < 3 {
		fmt.Print("masukan password ATM:")
		fmt.Scan(&password)

		if password != "2007" {
			kesempatan++
			fmt.Println("password salah. Silahkan coba lagi.")
		}
	}

	// jika password salah 3 kali
	if password !="2007" {
		fmt.Println("ATM terblokir. Program selesai.")
		return
	}

	fmt.Println("login berhasil")

	// menu atm

	for pilihan != 3 {
		fmt.Println("Memu ATM:")
				fmt.Println("1. Lihat saldo")
						fmt.Println("2. Tarik uang")
								fmt.Println("3. keluar")
										fmt.Print("Pilih:")
										fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Printf("Saldo Anda: Rp %d\n", saldo)
		} else if pilihan == 2 {
			var jumlah int

			fmt.Print("Masukan jumlah uang yang ingin ditarik: Rp")
			fmt.Scan(&jumlah)

			if jumlah > 0 && jumlah <= saldo {
				saldo -= jumlah
				fmt.Printf("anda berhasil menarik uang sebesar Rp %d\n", jumlah)
				fmt.Printf("saldo anda tersisa: Rp %d\n", saldo)
			} else {
				fmt.Println("saldo tidak mencukupi")
			}
		} else if pilihan == 3 {
			fmt.Println("keluar dari sistem. Terima kasih!")
		} else {
			fmt.Println("pilihan tidak valid.")
		}
	}
}