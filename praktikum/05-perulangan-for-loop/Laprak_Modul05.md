# <h1 align="center">Laporan Praktikum Modul 05 - Perulangan for</h1>
<p align="center">Amanullah Risky Maya Sekti - 109092600020 </p>

## Dasar Teori

### A. Perulangan ```for```

Perulangan ```for``` merupakan struktur kontrol dalam bahasa Go yang digunakan untuk menjalankan suatu blok kode secara berulang selama kondisi tertentu terpenuhi. Perulangan suatu blok kode secara berulang selama kondisi tertentu terpenuhi. Perulangan ini membantu menyelesaikan permasalahan yang membutuhkan proses berulang sehingga program menjadi lebih efisies dan terstruktur

### B. Komponen perulangan ```for```

Perulangan ```for``` dengan tiga komponen terdiri dari inisialisasi, kondisi, dan post statement. Inisialisasi digunakan untuk menentukan nilai awal variabel penghitung, kondisi digunakan untuk menentukan apakah perulanan dilanjutkan, sedangkan post statement digunakan untuk memperbarui nilai variabel setelah setiap iterasi. Selain itu, Go juga mendukung perulangan ```for``` dengan kondisi saja

### C. Penerapan perulangan dalam pemrograman

Perulangann dapat digunakan untuk memproses sejumlah data, melakukan perhitungan berulang, mencari nilai terbesar, menghitung jumlah tabungan, serta menentukan klasifikasi bilangan. Dalam bahasa Go, perulangan ```for``` dapat dikombinasikan dengan percabangan ```if``` untuk memeriksa kondisi tertentu selama proses berlangsung
<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. bioskop.go

```go
package main

import "fmt"

func main() {
	var t int
	fmt.Print("input t (jumlah transaksi): ")
	fmt.Scan(&t)

	totalPendapatan := 0

	for i := 1; i <= t; i++ {
		var jenisFilm, hari string
		var jumlahTiket int

		fmt.Printf("input transaksi ke-%d (jenis, hari, tiket): ", i)
		fmt.Scan(&jenisFilm, &hari, &jumlahTiket)

		hargaSatuan := 0

		if jenisFilm == "Reguler" {
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
				hargaSatuan = 40000
			} else {
				hargaSatuan = 50000
			}
		} else if jenisFilm == "3D" {
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
				hargaSatuan = 60000
			} else {
				hargaSatuan = 70000
			}
		} else if jenisFilm == "IMAX" {
			if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
				hargaSatuan = 80000
			} else {
				hargaSatuan = 90000
			}
		}

		totalHarga := hargaSatuan * jumlahTiket

		if jumlahTiket >= 5 {
			diskon := (totalHarga * 10) / 100
			totalHarga = totalHarga - diskon
		}

		fmt.Printf("output total harga transaksi ke-%d: %d\n", i, totalHarga)

		totalPendapatan = totalPendapatan + totalHarga
	}

	fmt.Println("output total pendapatan:", totalPendapatan)
}
```
#### Deskripsi

Program ini digunakan untuk menghitung harga tiket bioskop berdasarkan jenis film, hari pemutaran, dan jumlah tiket yang dibeli. Harga tiket ditentukan berdasarkan kategori Reguler, 3D, atau IMAX serta perbedaan harga pada hari kerja dan akhir pekan. Jika jumlah tiket yang dibeli minimal lima tiket, program memberikan diskon sebesar 10%. Program kemudian menampilkan total harga setiap transaksi dan menjumlahkan seluruh pendapatan dari transaksi yang dilakukan.


### 2. klasifikasi.go

```go
package main
import "fmt"
func main() {
	var a, b, n int
	fmt.Scan(&a, &b, &n)
	
	total := 0;
	tabunganHarian := a;

	for i := 1; i <= n; i++ {
		total += tabunganHarian
		tabunganHarian += b
	}
	fmt.Println(total)
}

```
#### Deskripsi

Program ini digunakan untuk menghitung total tabungan selama ```n``` hari dengan jumlah tabungan yang bertambah secara tetap setiap hari. Variabel ```a``` menyatakan tabungan awal, ```b``` menyatakan pertambahan tabungan harian, dan ```n``` menyatakan jumlah hari. Perulangan ```for``` digunakan untuk menjumlahkan tabungan setiap hari, kemudian meningkatkan nilai tabungan harian sebesar ```b```. Setelah perulangan selesai, program menampilkan total tabungan yang terkumpul.
### 3. Terbesar.go
```gopackage main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	
	var angka, max int
	fmt.Scan(&angka)
	max = angka
	for i := 2; i <= n; i++ {
		fmt.Scan(&angka)
		if angka > max {
			max = angka
		} 
		}

	
	fmt.Println(max)
}
```
#### Deskripsi
Program ini digunakan untuk mencari nilai terbesar dari sejumlah bilangan bulat yang dimasukkan pengguna. Bilangan pertama disimpan sebagai nilai maksimum sementara. Selanjutnya, program membaca bilangan berikutnya menggunakan perulangan ```for``` dan membandingkannya dengan nilai maksimum yang tersimpan. Jika ditemukan bilangan yang lebih besar, nilai maksimum diperbarui. Setelah seluruh bilangan diperiksa, program menampilkan nilai terbesar.
<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. prima.go

```go
package main
import "fmt"
func main () {
	var n int
	fmt.Scan(&n)

	var jumlahFaktor int

	for i :=1; i <=n; i++{
	if n%i == 0 {
		fmt.Println(i)
		jumlahFaktor++
	}
}

if jumlahFaktor == 2 {
	fmt.Println("ya")

} else {
	fmt.Println("tidak")

}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](unguided/Soal1/output.png)


#### Deskripsi

Program ini digunakan untuk menampilkan seluruh faktor dari bilangan positif ```n``` dan menentukan apakah bilangan tersebut prima. Program menghitung jumlah faktor menggunakan perulangan ```for```. Jika jumlah faktor tepat dua, program menampilkan ```ya```, jika tidak menampilkan ```tidak```.

### 2. ecct.go

```go

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

```

##### Output
![Screenshot Output Unguided](unguided/Soal2/output.png)    

#### Deskripsi

Program ini digunakan untuk menentukan kelayakan mahasiswa mendapatkan sertifikat ECCT berdasarkan enam nilai tes. Jika total nilai minimal 20, mahasiswa mendapat sertifikat. Program juga menghitung jumlah mahasiswa yang memperoleh sertifikat.

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan

Praktikum ini membahas penerapan struktur perulangan ```for``` dalam bahasa pemrograman Go untuk menyelesaikan berbagai permasalahan. Perulangan dapat digunakan untuk memproses banyak transaksi, menghitung tabungan, mencari nilai terbesar, menetukan bilangan prima, dan mengevaluasi nilai mahasiswa. Dengan menggabungkan perulangan dan percabangan ```if```, program dapat melakukan pengolahan data secara berulang berdasarkan kondisi tertentu. oleh karna itu, pemahaman mengenai perulangan ```for``` sangat penting untuk membuat program yang efisies, terstruktur, dan mudah dipahami

## Referensi
1. The Go Authors. A Tour of Go — For. Diakses pada 20 Oktober 2026 melalui https://go.dev/tour/flowcontrol/1



2. The Go Authors. Go Documentation. Diakses pada 20 Oktober 2026 melalui https://go.dev/doc/
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
