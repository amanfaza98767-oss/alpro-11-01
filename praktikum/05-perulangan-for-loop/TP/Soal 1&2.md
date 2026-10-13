1. Apa fungsi utama dari struktur perulangan for dalam program Go? Sebutkan dan jelaskan tiga komponen pada perulangan for (inisialisasi, kondisi, dan post statement)! 
2. Bahasa Go tidak memiliki perulangan while dan do-while. Jelaskan perbedaan antara for dengan tiga komponen dan for yang hanya menggunakan kondisi! Berikan contoh kode masing-masing. 

Jawaban

1. Fungsi utama for adalah untuk menjalankan blok kode secara berulang selama kondisi tertentu terpenuhi
    1. Inisialisasi: Menentukan nilai awal variabel yang digunakan dalam perulangan
    2. Kondisi: Menentukan apakah perulangan akan dilanjutkan. Jika kondisi bernilai true, perulangan berjalan; jika kondisi false, perulangan berhenti
    3. Post statement: Menjalankan perubahan variabel setelah setiap iterasi, misalnya menambah nilai variabel dengan i++

2. A. For dengan tiga komponen
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