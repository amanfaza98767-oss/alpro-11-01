# <h1 align="center">Laporan Praktikum Modul 03 - Variable dan Operator</h1>
<p align="center">Amanullah Risky Maya Sekti0 - 109092600020</p>

## Dasar Teori

### A. Variable
Variable adalah tempat untuk menyimpan data yang digunakan selama program berjalan. Dalam bahasa Go, variable dapat dideklarasikan menggunakan kata kunci ```var``` atau menggunakan deklarasi singkat ```:=```. Setiap variable memiliki tipe data tertentu, seperti ```int``` untuk bilangan bulat dan ```float64``` untuk bilangan desimal.

### B. Operator
Operator digunakan untuk melakukan operasi terhadap data atau variabel. Dalam bahasa Go terdapat beberapa operator aritmatika, seperti ```+``` untuk penjumlahan, ```-``` untuk pengurangan, ```*``` untuk perkalian, ```/``` untuk pembagian, dan ```%``` untuk mendapatkan sisa hasil pembagian.

### C. Tipe Data dan Deklarasi Variabel di Go
Tipe data menentukan jenis nilai yang dapat disimpan oleh sebuah variabel. Beberapa tipe data dasar yang digunakan dalam praktikum adalah. 

#### 1. Integer (```int)
Digunakan untuk menyimpan bilanagn bulat, seperti jumlah hari atau nominal uang.

#### 2. Float (```float64```)
Digunakan untuk menyimpan bilanagn desimal, seperti hasil konversi suhu.
#### 3. Deklarasi variable
Variabel dapat dibuat dengan ```var```, contohnya:
```go
 var totalHari int
```
Go juga mendukung deklarasi singkat:
```go 
suhu := 25,5
```

### D. Operasi Pembagian dan Modulus
Pembagian bilangan bulat menggunakan operator ```/``` menghasilkan nilai hasil bagi, sedangkan operator ```%``` menghasilkan sisa pembagian. Kedua operator tersebut digunakan dalam program pencacahan uang dan pencacahan hari untuk memisahkan suatu nilai menjadi beberapa satuan

Contohnya:
```go
totalHari := 365
tahun := totalHari / 360
sisa := totalHari % 360

Pada contoh tersebut, ```/``` digunakan untuk memperoleh jumlah tahun, sedangkan ```%``` digunakan untuk memperoleh sisa hari

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1.  cacah uang

```go
package main
import ("fmt")

func main() {
	var totalUang int

	//membaca masukan nominal uang
	fmt.Scan(&totalUang)

	//menghitung lembar puluhan ribu
	sepuluhribu := totalUang / 10000
	sisa := totalUang % 10000

	//menghitung lembar lima ribu
	limaribu := sisa / 5000
	sisa = sisa % 5000

	//menghitung lembar  seriburibu
	seribu := sisa / 1000
	
	//menampilkan hasil keluaran sesuai format
	fmt.Printf("%d %d %d\n", sepuluhribu, limaribu, seribu)
}


```
#### Deskripsi
program ini dibuat untuk mengetahui banyaknya becahan 10 ribu, 5 ribu, dan seribu

### 2. Konversi suhu C -> K

```go
package main

import "fmt"

func main() {
	var  celcius float64
	fmt.Print("Masukkan suhu dalam Celcius: ")
	fmt.Scan(&celcius)
	var kelvin float64 = celcius + 273
	fmt.Println("Suhu dalam Kelvin:", kelvin)
}

```
#### Deskripsi
program ini dibuat untunk mengubah suhu dari celcius menjadi kelvin

### 2. Pertukaran angka

```go
package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)

	temp := x
	x = z
	z = y
	y = temp

	fmt.Println(x, y, z)
}


```
#### Deskripsi
program ini digunakan untuk mengganti posisi angka di dalam variable  X, Y, dan Z
<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. Cacah Hari

```go
package main
import "fmt"

func main() {
var totalHari int

fmt.Scan(&totalHari)

tahun := totalHari / 360
sisa := totalHari % 360

bulan := sisa / 30
sisa = sisa % 30

minggu := sisa / 7
sisa = sisa % 7

fmt.Printf("%d %d %d %d\n", tahun, bulan, minggu, sisa)
}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](unguided/cacah%20hari/output.png)


#### Deskripsi
program ini dibuat untuk mengetahui banyaknya tahun, bulan, dan minggu dengan cara menginput banyaknya hari

### 2. Konversi suhu C -> R

```go
package main
import "fmt"

func main() {
	var celcius float64
	fmt.Print("Masukkan suhu dalam Celcius: ")
	fmt.Scan(&celcius)
	var reamur float64 = celcius * 4 / 5
	fmt.Println("Suhu dalam Reamur:", reamur)
}
```

##### Output
![Screenshot Output Unguided](unguided/konversi%20C%20ke%20R/output.png)

#### Deskripsi
program ini di gunakan untuk mengkonversi suhu dari celcius ke reamur

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
Praktikum modul 03 tentang Variable dan Operator memberikan pemahaman mengenai penggunaan variabel, tipe data, serta operator dalam bahasa pemrograman Go. Melalui program yang dibuat, dapat diterapkan operator aritmatika seperti pembagian (```/```), modulus (```%```), perkalian (```*```), dan penjumlahan (```+```) untuk menyelesaikan berbagai permasalahan.

Program yang dibuat meliputi pencacahan uang, konversi suhu, pertukaran nilai variabel, pencacahan hari, dan konversi suhu Celcius ke Reamur. Dari praktikum ini dapat disimpulkan bahwa pemahaman mengenai variabel dan operator merupakan dasar penting dalam membuat program karena keduanya digunakan untuk menyimpan, mengolah, dan menghasilkan data sesuai kebutuhan.

## Referensi
1. Donovan, A. A. A., & Kernighan, B. W. (2015). The Go Programming Language. Addison-Wesley Professional. Diakses pada tanggal 4 oktober 2026 melalui https://books.google.co.id/books?id=SJHvCgAAQBAJ&printsec=frontcover&hl=fr&source=gbs_ge_summary_r&cad=0#v=onepage&q&f=false
2. The Go Authors. (2026). The Go Programming Language Specification. Diakses pada tanggal 4 oktober 2026 melalui [Go.dev – The Go Programming Language](https://go.dev/ref/spec?utm_source)
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
