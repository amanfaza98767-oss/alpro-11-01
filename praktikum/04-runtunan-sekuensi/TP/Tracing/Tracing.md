Soal 2 – Tracing: 
Evaluasi Pernyataan Kondisi Berikut adalah kode dalam bahasa Go yang berisi beberapa pernyataan kondisi. Tugas Anda adalah melacak alur eksekusi program dan menentukan nilai akhir dari setiap variabel setelah program selesai dijalankan. Selain itu, tuliskan juga output yang dihasilkan oleh program. 
Pertanyaan: 
1. Berapa nilai akhir dari variabel result setelah semua pernyataan kondisi dieksekusi? 25
2. Apa output yang dihasilkan oleh program? Output akhir: 25
3. Tuliskan langkah-langkah alur eksekusi program berdasarkan kondisi yang diberikan: 
◦ Kondisi 1: Apakah kondisi x > 5 benar? Benar. Jika ya, apa yang terjadi selanjutnya? 
```
result = x + y
       = 10 + 5
       = 15
```
◦ Kondisi 2: Apakah kondisi z > 10 && x == 10 benar? Benar. Bagaimana hal ini memengaruhi nilai result? 
```
result += z
result = 15 + 15
result = 30
```
◦ Kondisi 3: Apakah salah satu dari kondisi x == 10 || y > 10 benar? Benar. Apa yang terjadi? 
```
result += 5
result = 30 + 5
result = 35
```
◦ Kondisi 4: Bagaimana kondisi !(x < 15 && y < 10) dievaluasi? Salah. Apa dampaknya pada result? 
```
!(true) = false
```
jadi kondisi if adalah false, sehingga masuk ke else
```
result -= 10
result = 35 - 10
result = 25
```