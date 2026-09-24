package main
import ("fmt")

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	tambah := a + b
	kurang := a - b
	kali := a * b
	bagi := a / b
	sisa := a % b

	fmt.Println("Hasil Penjumlahan:", tambah)
	fmt.Println("Hasil Pengurangan:", kurang)
	fmt.Println("Hasil Perkalian:", kali)
	fmt.Println("Hasil Pembagian:", bagi)
	fmt.Println("Hasil Sisa Bagi:", sisa)
}
