# <h1 align="center">Tugas Pendahuluan Modul 04 - Runtutan/Sekuensi</h1>
<p align="center">Amanullah Risky Maya Sekti - 109092600020</p>

### 1. Evaluasi Ekspresi

```go
package main

import "fmt"

func main() {
	intNum := 5
	intOther := 10
	var sngNum float64 = -3

	fmt.Println("1.", intNum > 5)

	fmt.Println("2.", intNum >= 5 && intOther < 11)

	fmt.Println("3.", sngNum != -1 || intOther < 0)

	fmt.Println("4.", !(intNum > 3) || intNum <= 5)

	fmt.Println("5.", !(intOther >= intNum))

	fmt.Println("6.", 0-sngNum > 0)

	fmt.Println("7.", 4/2 == intOther/intNum)

	fmt.Println("8.", intOther%2 == 0)

	fmt.Println("9.", intOther+2*intNum != 30 || !(sngNum > 0))

	fmt.Println("10.", intOther > 0 && intNum > 0 || sngNum > 0)

	fmt.Println("11.", sngNum > 0 || (intNum >= 0 && -1*intOther == -10))

	fmt.Println("12.", intNum == 5)

	fmt.Println("13.", intNum > 0 || (sngNum <= 0 && intOther == 13))

	fmt.Println("14.", !(!(!(!(intNum > 0)))))
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](tp/evaluasi%20ekspresi/Output.png)


#### Deskripsi
Pada praktikum ini dilakukan evaluasi terhadap beberapa ekspresi dalam bahasa Go yang menggunakan operator perbandingan, operator logika, operator aritmatika, dan operator NOT. Setiap ekspresi dihitung berdasarkan nilai variabel yang telah ditentukan, kemudian hasilnya ditampilkan dalam bentuk nilai boolean (true atau false). Dari proses tersebut diperoleh hasil evaluasi untuk 14 ekspresi sesuai dengan aturan prioritas operator dalam bahasa Go.


### 2. Menentukan jumlah hari

```go
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
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](tp/Menentukan%20Hari/output.png)


#### Deskripsi
Pada praktikum ini dibuat program untuk menentukan jumlah hari dalam suatu bulan berdasarkan input tahun dan nama bulan. Program menggunakan struktur switch untuk menentukan jumlah hari pada setiap bulan serta percabangan if untuk menentukan apakah bulan Februari berada pada tahun kabisat. Program menerapkan aturan tahun kabisat, yaitu tahun yang habis dibagi 4 dan tidak habis dibagi 100, atau habis dibagi 400. Hasil program menampilkan jumlah hari sesuai dengan tahun dan bulan yang dimasukkan.


### 3 Tracing
```go
package main
import ("fmt")

func main() {
	x := 10
	y := 5
	z := 15
	result := 0

	if x > 5 {
		if y < 10 {
			result = x + y
		} else {
			result = x - y
		}
	}
	if z > 10 && x == 10 {
		result += z
	} else {
		result = z - x
	}
	if x == 10 || y > 10 {
		result += 5
	} else if y == 5 && z > 10 {
		result -= 5
	} else {
		result *= 2
	}
	if ! (x< 15 && y < 10) {
		result += 10
	} else {
		result -= 10
	}
	fmt.Println("Final Result:", result)
}
	
```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](tp/Tracing/output.png)


#### Deskripsi
Pada praktikum ini dilakukan tracing terhadap alur eksekusi program menggunakan beberapa struktur percabangan if, else if, dan else. Nilai setiap variabel diperiksa berdasarkan kondisi yang diberikan, kemudian nilai variabel result diperbarui sesuai dengan operasi yang dijalankan. Dengan melakukan tracing secara berurutan, dapat diketahui perubahan nilai result hingga program selesai dan menghasilkan nilai akhir yang ditampilkan sebagai Final Result.



### 4 Lampu lalu lintas
``` go
package main
import ("fmt")

func main() {
	var warna string
	fmt.Print("Masukkan warna lampu lalu lintas (merah, kuning, hijau): ")
	fmt.Scan(&warna)

	switch warna {
	case "merah":
		fmt.Println("Lampu merah: Berhenti")
	case "kuning":
		fmt.Println("Lampu kuning: Hati-hati")
	case "hijau":
		fmt.Println("Lampu hijau: Jalan")
	default:
		fmt.Println("Warna lampu tidak valid")
	}
}
```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](tp/Lampu%20lalu%20lintas/output.png)


#### Deskripsi
Pada praktikum ini dibuat program untuk menentukan tindakan berdasarkan warna lampu lalu lintas yang dimasukkan oleh pengguna. Program menggunakan struktur switch untuk memeriksa input warna, yaitu merah, kuning, dan hijau. Setiap warna memiliki keluaran yang berbeda, sedangkan input selain ketiga warna tersebut akan menghasilkan pesan bahwa warna lampu tidak valid. Program berhasil menampilkan tindakan yang sesuai berdasarkan warna yang dimasukkan.


## Kesimpulan
Berdasarkan praktikum yang telah dilakukan, dapat disimpulkan bahwa bahasa Go menyediakan berbagai struktur dan operator untuk mengatur alur eksekusi program. Operator perbandingan dan logika dapat digunakan untuk mengevaluasi suatu kondisi, sedangkan struktur switch dan if-else dapat digunakan untuk menentukan tindakan berdasarkan kondisi tertentu. Melalui proses tracing, perubahan nilai variabel selama program berjalan juga dapat diketahui secara sistematis. Dengan demikian, praktikum ini membantu memahami penerapan ekspresi, percabangan, switch, dan penelusuran alur program dalam bahasa Go.
