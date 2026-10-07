Soal 1 – Evaluasi Ekspresi Kontrol dalam Go Diberikan segmen kode berikut, di mana setiap ekspresi kontrol yang terdaftar di bawah dapat digunakan sebagai kondisi pada pernyataan if: 
``` go
intNum := 5 
intOther := 10 
var sngNum float64 = -3  
if ____________________ {     fmt.Println("Beep") } 
``` 
Untuk setiap latihan di bawah ini, tentukan apakah ekspresi kontrol tersebut menghasilkan true atau false. Tuliskan hasilnya dengan kata "true" atau "false". 
No true/false Kondisi 
1. False intNum > 5 
2. True intNum >= 5 && intOther < 11 
3. True sngNum != -1 || intOther < 0 
4. True !(intNum > 3) || intNum <= 5 
5. False !(intOther >= intNum) 
6. True 0 - sngNum > 0 
7. True 4 / 2 == intOther / intNum 
9. True intOther + 2 * intNum != 30 || !(sngNum > 0) 
8. True intOther % 2 == 0 
10. True intOther > 0 && intNum > 0 || sngNum > 0 
11. True sngNum > 0 || (intNum >= 0 && -1 * intOther == -10) 
12. True intNum == 5 
13. True intNum > 0 || (sngNum <= 0 && intOther == 13) 
14. True !(!(!(!(intNum > 0)))) 