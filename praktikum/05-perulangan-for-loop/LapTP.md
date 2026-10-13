# <h1 align="center">Tugas Pendahuluan Modul 05  - Perulangan for</h1>
<p align="center">Amanullah Risky Maya Sekti - 109092600020</p>

### 1. Apa fungsi utama dari struktur perulangan for dalam program Go? Sebutkan dan jelaskan tiga komponen pada perulangan for (inisialisasi, kondisi, dan post statement)! 

1. Fungsi utama for adalah untuk menjalankan blok kode secara berulang selama kondisi tertentu terpenuhi
    1. Inisialisasi: Menentukan nilai awal variabel yang digunakan dalam perulangan
    2. Kondisi: Menentukan apakah perulangan akan dilanjutkan. Jika kondisi bernilai true, perulangan berjalan; jika kondisi false, perulangan berhenti
    3. Post statement: Menjalankan perubahan variabel setelah setiap iterasi, misalnya menambah nilai variabel dengan i++



### 2.  Bahasa Go tidak memiliki perulangan while dan do-while. Jelaskan perbedaan antara for dengan tiga komponen dan for yang hanya menggunakan kondisi! Berikan contoh kode masing-masing. 

A. For dengan tiga komponen
    Perulangan ini memiliki inisialisasi, kondisi, dan post statement. Bentuknya mirip dengan perulangan for pada bahasa C atau java
    ```go
    package main
    import "fmt"
    func() {
        for i := 1; i <= 5; i++ {
            fmt.Println(i)
        }
    }
    ```

    B. For dengan kondisi saja
    Perulangan ini hanya menggunakan kondisi tanpa inisialisasi dan post statement pada bagian for. Bentuk ini mirip dengan while pada bahasa lain
    ```go
    package main
    import "fmt"
    func () {
        i := 1
        for i <= 5 {
            fmt.Println(i)
            i++
        }
    }
    ```

    Perulangan for dengan tiga komponen mengatur inisialisai, kondisi, dan perubahan variabel dalam satu baris. Sementara itu, for dengan kondisi saja hanya menuliskan kondisi pada bagian perulangan, sehingga inisialisasi dan perubahan variabel dilakukan secara terpisah. Keduanya dapat digunakan untuk menjalankan kode berulang di Go

### 3 Ganjil.go
```go
package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	jumlah := 0
	for i := 1; i <= n; i += 2 {
		fmt.Println(i, "")
		jumlah += i
	}
	fmt.Println("Jumlah bilangan ganjil dari 1 sampai", n, "adalah", jumlah)
}
	
```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](tp/soal3/output.png)


#### Deskripsi
Program ini digunakan untuk menampilkan semua bilangan ganjil dari 1 hingga n menggunakan perulangan ```for```. Setiap bilangan ganjil ditampilkan dan dijumlahkan ke dalam variabel ```jumlah```. Setelah perulangan selesai, program menampilkan total penjumlahan seluruh bilangan ganjil tersebut.


### 4 Perkalian
``` go
package main
import "fmt"
func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	for i := 1; i <= int(b); i++ {
		fmt.Println(a, "x", i, "=", a*float64(i))
	} 
	fmt.Println("Perkalian", a, "sampai", b, "adalah", a*b)
}
```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](tp/soal4/output.png)


#### Deskripsi
Program ini digunakan untuk menampilkan tabel perkalian dari bilangan a mulai dari a × 1 hingga a × b. Program menerima dua input bertipe float64, kemudian menggunakan perulangan for untuk menampilkan hasil perkalian pada setiap iterasi. Setelah perulangan selesai, program menampilkan hasil perkalian a × b sebagai hasil akhir.


## Kesimpulan
Perulangan ```for``` dalam bahasa Go digunakan untuk menjalankan instruksi secara berulang sesuai kondisi yang ditentukan. Perulangan ini dapat digunakan untuk menmpilkan bilangan ganjil sekaligus menghitung jumlahnya, serta membuat tabel perkalian. Dengan memahami perulangan ```for```, program dapat dibuat lebih efisien tanpa harus menuliskan instruksi yang sama berulang kali.
