# <h1 align="center">Laporan Praktikum Modul 04 - Runtutan dan Sekuensi</h1>
<p align="center">Amanullah Risky Maya Sekti - 109092600020 </p>

## Dasar Teori

### A. Runtutan atau Sekuensi
Runtutan atau sekuensi merupakan struktur dasar algoritma yang menjelaskan intruksi satu demi satu sesuai urutan penulisannya. Setiap intruksi dikerjakan sesuai alur program, sehingga urutan operasi dapat memengaruhi hasil akhir

### B. Struktur Percabangan switch
Pernyataan switch  digunakan untuk memilih tindakan berdasarkan nilai atau kondisi tertentu. Dalam Go, switch dapat ditulis tanpa ekspresi sehingga setiap case berisi kondisi boolean. Program menjalankan blok case pertama yang kondisinya bernilai benar>

### C. Bufio
bufio adalah package dalam bahasa pemrograman Go yang digunakan untuk membaca dan menulis data menggunakan buffer agar proses input dan output lebih efisien.
<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. grade.go

```go
package main
import (
"bufio"
 "fmt"
 "os"
)

func main() {
	var nama string
	var nilai float64
	var grade string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	nama = scanner.Text()

	fmt.Scan(&nilai)

	if nilai >= 90 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 80 && nilai < 90 {
		grade = "B"
	} else if nilai >= 70 && nilai < 80 {
		grade = "C"
	} else if nilai >= 60 && nilai < 70 {
		grade = "D"
	} else if nilai >= 0 && nilai < 60 {
		grade = "F"
	}
	fmt.Printf("%s mendapatkan nilai %s", nama, grade)
}
```
#### Deskripsi
untuk menentukan grade dengan cara menginput nama dan nilai

### 2. penilaian.go

```go
package main

import (
	"bufio" // Digunakan untuk membaca masukan string yang memiliki spasi (seperti nama lengkap)
	"fmt"
	"os" // Menyediakan akses ke sistem operasi, dalam hal ini os.Stdin yang merepresentasikan keyboard
)

func main() {
	// Pendefinisian variabel secara eksplisit beserta tipe datanya
	var pilihan int
	var nama string
	var nilai float64
	var grade string

	// Menampilkan cetakan Menu ke layar
	fmt.Println("============= Menu =============")
	fmt.Println("1. Sistem penilaian")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih opsi: ")

	// Membaca masukan pilihan (angka).
	// Kita menggunakan Scanln agar saat user menekan 'Enter', karakter enter tersebut
	// ikut diolah/dibersihkan dan tidak melompat (mengganggu) masukan nama di bawahnya.
	fmt.Scanln(&pilihan)

	// Percabangan/Sekuensi bersyarat berdasarkan input pengguna
	if pilihan == 1 {
		fmt.Print("Masukkan nama siswa: ")

		// Membuat alat pembaca (scanner) baru yang mengambil masukan dari keyboard
		scanner := bufio.NewScanner(os.Stdin)

		// Proses membaca masukan dari pengguna sampai tombol Enter ditekan
		scanner.Scan()

		// Mengambil teks (nama) yang baru saja dibaca dan menyimpannya ke variabel
		nama = scanner.Text()

		fmt.Print("Masukkan nilai siswa: ")
		fmt.Scanln(&nilai)

		// Menentukan grade nilai
		if nilai >= 90 && nilai <= 100 {
			grade = "A"
		} else if nilai >= 80 && nilai < 90 {
			grade = "B"
		} else if nilai >= 70 && nilai < 80 {
			grade = "C"
		} else if nilai >= 60 && nilai < 70 {
			grade = "D"
		} else {
			grade = "F"
		}

		// %s adalah format (placeholder) untuk mencetak data bertipe string
		// %s pertama akan digantikan oleh isi variabel 'nama', %s kedua oleh 'grade'
		fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)

	} else if pilihan == 0 {
		// Dijalankan jika pengguna mengetik 0
		fmt.Println("Keluar dari program.")

	} else {
		// Dijalankan jika pengguna mengetik angka selain 1 dan 0
		fmt.Println("Pilihan tidak valid.")
	}
}
```
#### Deskripsi
program ini sama seperti program sebelumnya hanya saja di tambahkan tampilan menu
### 3. klasifikasi.go

```go
package main

import (
	"fmt" // Hanya memerlukan fmt untuk keperluan input dan output
)

func main() {
	// Pendefinisian variabel secara eksplisit beserta tipe datanya
	var usia int
	var gaji int
	var keterangan string

	// Membaca dua masukan berupa angka dari pengguna (usia dan gaji)
	// fmt.Scan otomatis memisahkan input berdasarkan spasi atau baris baru (enter)
	fmt.Scan(&usia)
	fmt.Scan(&gaji)

	// Menggunakan switch tanpa ekspresi.
	// Cara kerjanya sama persis seperti deretan if - else if.
	// Program akan mengecek dari atas ke bawah, dan menjalankan case pertama yang bernilai benar (true).
	switch {
	case usia < 18:
		keterangan = "Masih sekolah"
	case usia >= 18 && usia <= 25 && gaji >= 50:
		keterangan = "Muda sukses"
	case usia >= 18 && usia <= 25 && gaji < 50:
		keterangan = "Masih belajar hidup"
	case usia >= 26 && usia <= 40 && gaji >= 100:
		keterangan = "Pekerja mapan"
	case usia >= 26 && usia <= 40 && gaji < 100:
		keterangan = "Perlu perbaikan karier"
	case usia > 40 && gaji >= 150:
		keterangan = "Profesional berpengalaman"
	case usia > 40 && gaji < 150:
		keterangan = "Perlu evaluasi finansial"
	}

	// Menampilkan hasil klasifikasi ke layar
	fmt.Println(keterangan)
}
```
#### Deskripsi
program ini digunakan untuk megelompokan seseorang berdasarkan usia dan gaji dengan cara memasukan usia dan gaji, dengan menggunakan stuktur switch

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. grade.go

```go
package main
import ("bufio"
"fmt"
"os"

)

func main() {
	var nama string
	var nilai float64
	var grade string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nama = scanner.Text()
	fmt.Scan(&nilai)
	switch {
	case nilai >= 90 && nilai <= 100:
		grade = "A"
	case nilai >= 80 && nilai < 90:
		grade = "B"
	case nilai >= 70 && nilai < 80:
		grade = "C"
	case nilai >= 60 && nilai < 70:
		grade = "D"
	case nilai >= 0 && nilai < 60:
		grade = "F"
		case nilai < 0 || nilai > 100:
		fmt.Println("Nilai tidak valid")
		return
	}


fmt.Printf("%s mendapatkan nilai %s", nama, grade)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](unguided/soal%201/output.png)


#### Deskripsi
untuk menentukan nilai dari sebuah mahasiswa/siswa dengan cara memasukan nilainya dan outputnya berupa A, B, C, D, dan F dengan menggunakan struktur switch

### 2. pajak.go

```go
package main
import (
"fmt"
)

func main() {
	var pajak float64
	var gaji float64

	fmt.Scan(&gaji)
	
if gaji <= 50 {
		pajak = 0.05 * gaji
	} else if gaji > 50 && gaji <= 100 {
		pajak = 0.05 * 50
		pajak += 0.1 * (gaji - 50)
	} else if gaji > 100 && gaji <= 200 {
		pajak = 0.05 * 50
		pajak += 0.1 * 50
		pajak += 0.15 * (gaji - 100)
	} else if gaji > 200 {
		pajak = 0.05 * 50
		pajak += 0.1 * 50
		pajak += 0.15 * 100
		pajak += 0.2 * (gaji - 200)
	}
	fmt.Printf("Pajak yang harus dibayarkan: %.2f", pajak)

}
```

##### Output
![Screenshot Output Unguided](unguided/soal%202/output.png)

#### Deskripsi
untuk menghitung pajak dengan cara memberikan informasi berupa jumlah peghasilan dan di hitung dengan cara persentase dikali jumlah penghasilan

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
program Go menjalankan instruksi secara berurutan serta dapat menggunakan percabangan untuk menentukan tindakan berdasarkan kondisi tertentu. Struktur if-else dan switch dapat diterapkan untuk menentukan grade nilai, membuat menu penilaian, mengklasifikasikan usia dan gaji, serta menghitung pajak berdasarkan penghasilan. Praktikum ini membantu memahami penggunaan variabel, input-output, operator perbandingan, operator logika, dan percabangan dalam menyelesaikan permasalahan menggunakan algoritma pemrograman.

## Referensi
1. The Go Authors. (n.d.). Go Documentation: Package bufio. Diakses pada 11 Oktober 2026 melalui https://pkg.go.dev/bufio
2. Go by Example. (n.d.). Switch. Diakses pada 11 Oktober 2026 melalui https://gobyexample.com/switch
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
