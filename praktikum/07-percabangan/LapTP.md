# <h1 align="center">Tugas Pendahuluan Modul 06 - PERULANGAN WHILE</h1>
<p align="center">Amanullah Risky Maya Sekti - 109092600020</p>

### 1. Lulus

```go
package main
import "fmt"
func main() {
 var nilai int

 for {
	fmt.Print("masukan nilai anda: ")
	fmt.Scan(&nilai)
	if nilai == -1 {
		fmt.Println("program selesai")
		break
	}
	if nilai >= 60 {
		fmt.Println("anda lulus ujian")
	} else {
		fmt.Println("anda tidak lulus ujian")
	}
 }
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](tp/Soal1/output.png)


#### Deskripsi

Program menentukan kelulusan berdasarkan nilai yang dimasukkan. Jika nilai minimal 60, pengguna dinyatakn lulus. Program berhenti ketika nilai ``` -1 ``` dimasukkan.

### 2. Deret

```go
package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	i := 1
	jumlah := 0

	for i <= n {
		jumlah += i
		i++
	}
	fmt.Println(jumlah)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](tp/Soal2/output.png)


#### Deskripsi

Program menghitung jumlah bilangan dari 1 sampai n menggunakan perulangan dan menampilkan hasilnya.

### 3 duaterendah
```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var nama string
	var nilai int

	var namaTerendah, namaKedua string
	var nilaiTerendah, nilaiKedua int

	i := 1
	for i <= n {
		fmt.Scan(&nama, &nilai)

		if i == 1 {
			namaTerendah = nama
			nilaiTerendah = nilai
		} else if i == 2 {
			if nilai < nilaiTerendah {
				namaKedua = namaTerendah
				nilaiKedua = nilaiTerendah

				namaTerendah = nama
				nilaiTerendah = nilai
			} else {
				namaKedua = nama
				nilaiKedua = nilai
			}
		} else {
			if nilai < nilaiTerendah {
				namaKedua = namaTerendah
				nilaiKedua = nilaiTerendah

				namaTerendah = nama
				nilaiTerendah = nilai
			} else if nilai < nilaiKedua {
				namaKedua = nama
				nilaiKedua = nilai
			}
		}

		i++
	}

	fmt.Println(namaTerendah)
	fmt.Println(namaKedua)
}
```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](tp/Soal3/output.png)


#### Deskripsi

Program mencari dan menampilkan nama dua mahasiswa dengan nilai terendah dari seluruh data yang dimasukan.

### 4 Biaya
``` go
package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	var biaya float64 = 10000000
	i := 1

	for i < n {
		biaya = biaya * 1.06
		i++
	}

	fmt.Printf("%.0f\n", biaya)
}
```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](tp/Soal4/output.png)


#### Deskripsi

Program menghitung biaya kuliah tahun ke-n dengan biaya Rp.10.000.000 dan kenaikan sebesar 6% setiap tahun. 

## Kesimpulan

Perulangan ```for``` dalam Go dapat digunakan seperti ```while``` untuk menjalankan intruksi selama kondisi terpenuhi. Perulangan membantu menyelesaikan berbagai permasalahan, seperti memeriksa kelulusan, menghitung deret, mencari nilai terendah, dan menghitung kenaikan biaya kuliah secara efisien