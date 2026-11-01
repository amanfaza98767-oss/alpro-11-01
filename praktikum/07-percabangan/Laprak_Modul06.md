# <h1 align="center">Laporan Praktikum Modul 06 - Perulangan while loop</h1>
<p align="center">Amanullah Risky Maya Sekti - 109092600020 </p>

## Dasar Teori

### A. Pengertian Perulangan While
Perulangan while adalah struktur perulangan yang menjalankan instruksi selama kondisi bernilai benar. Perulangan akan berhenti ketika kondisi bernilai salah. Dalam bahasa Go, perulangan while dapat dibuat menggunakan for yang hanya memiliki kondisi.

### B. Cara Kerja Perulangan
Perulangan dimulai dengan memeriksa kondisi. Jika kondisi benar, instruksi di dalam perulangan dijalankan. Setelah itu, kondisi diperiksa kembali hingga bernilai salah. Variabel yang digunakan perlu diperbarui agar perulangan tidak berjalan tanpa henti.


### C. Penerapan Perulangan
Perulangan dapat digunakan untuk berbagai kebutuhan, seperti menghitung bilangan, memproses karakter dalam kalimat, dan menjalankan menu program secara berulang. Penggunaan perulangan membuat program lebih efisien karena instruksi tidak perlu ditulis berkali-kali.

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. terkecil.go

```go
package main
import "fmt"
func main() {
	var batas int
    fmt.Scan(&batas)

    n := 1
    for n*n*n <= batas {
        n++
    }

    fmt.Println(n)
}
```
#### Deskripsi
program menerima sebuah bilangan sebagai batas, kemudian mencari bilangan bulat positif terkecil yang nilai pangkat tiganya lebih besar dari batas tersebut. Variabel n dimulai dari 1 dan terus bertambah selama nilai n³ masih kurang dari atau sama dengan batass. Setelah kondisi tidak terpenuhi, program menampilkan nilai n.


### 2. terbesar.go

```go
package main
import "fmt"
func main() {
	var batas int
	 fmt.Scan(&batas)

    n := 1
    for n*n < batas {
        n++
    }

    fmt.Println(n - 1)
}

```
#### Deskripsi
Program ini menerima sebuah bilangan sebagai batas, kemudian mencari bilangan bulat positif terbesar yang nilai kuadratnya lebih kecil dari batas tersebut. Variabel n dimulai dari 1 dan bertambah selama nilai n² masih lebih kecil dari batas. Setelah perulangan selesai, program menampilkan n - 1 sebagai hasil.

### 3. vokal.go
```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    kalimat, _ := reader.ReadString('\n')

    i := 0
    jumlahVokal := 0

    for i < len(kalimat) {
        ch := strings.ToLower(string(kalimat[i]))
        if ch == "a" || ch == "i" || ch == "u" || ch == "e" || ch == "o" {
            jumlahVokal++
        }
        i++
    }

    fmt.Println(jumlahVokal)
}
```
#### Deskripsi
Program ini membaca sebuah kalimat, kemudian menghitung jumlah huruf vokal yang terdapat di dalamnya, yaitu a, i, u, e, o. Setiap karakter diubah menjadi huruf kecil agar huruf kapital juga dapat dihitung. Perulangan dilakukan hingga seluruh karakter kalimat diperiksa, kemudian jumlah huruf vokal ditampilkan.

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. atm.go

```go
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
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](unguided/soal1/output.png)


#### Deskripsi
Program ATM ini menggunakan perulangan untuk proses login dan menu transaksi. Pengguna memiliki maksimal tiga kali kesempatan memasukkan password. Jika berhasil login, pengguna bisa melihat saldo, menarik uang, atau keluar dari program. Saat menarik uang, program akan memeriksa apakah jumlah penarikan sesuai dengan saldo yang tersedia


## Kesimpulan
Pada praktikum Modul 06, saya mempelajari penggunaan perulangan while untuk menjalankan program secara berulang selama kondisi terpenuhi. Perulangan ini dapat digunakan untuk perhitungan, pengolahan data, dan pembuatan menu ATM agar program lebih efisien dan mudah digunakan.

## Referensi
1. The Go Authors. A Tour of Go — For. Diakses pada 28 Oktober 2026 melalui https://go.dev/tour/flowcontrol/1

2. The Go Authors. Go Documentation. Diakses pada 28 Oktober 2026 melalui https://go.dev/doc/