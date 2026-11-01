Jelaskan perbedaan antara if, if-else, else if, dan switch pada bahasa pemrograman Go! Sebutkan juga dua perbedaan switch pada Go dibandingkan dengan switch pada Java. 

jawab:
- if digunakan untuk menjalankan kode jika suatu kondisi bernilai benar (true).
- if-else digunakan untuk memilih satu dari dua kemungkinan berdasarkan kondisi. Jika kondisi if salah, blok else dijalankan.
- else if digunakan untuk memeriksa beberapa kondisi secara berurutan. Kondisi pertama yang benar akan dijalankan.
- switch digunakan untuk memilih tindakan berdasarkan nilai atau kondisi tertentu. Setiap pilihan ditulis dalam blok case, sedangkan default dijalankan jika tidak ada pilihan yang cocok.

- Perilaku case: Pada Go, setiap case secara otomatis berhenti setelah dijalankan sehingga tidak perlu menggunakan break. Pada Java, switch tradisional memerlukan break untuk mencegah eksekusi berlanjut ke case berikutnya.
- Ekspresi switch: Go memungkinkan switch tanpa ekspresi, sehingga setiap case dapat berisi kondisi boolean. Java juga mendukung switch dengan berbagai bentuk ekspresi pada versi modern, tetapi penggunaan switch tanpa ekspresi untuk memeriksa kondisi boolean secara langsung bukan sintaks standar Java.