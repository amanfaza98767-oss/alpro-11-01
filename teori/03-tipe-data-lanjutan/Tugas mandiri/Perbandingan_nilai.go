package main
import ("fmt")

func main() {
	var a, b int

	fmt.Scanf("%d", &a)

	fmt.Scanf("%d", &b)
	fmt.Println("Hasil perbandingan a > b:", a > b)
	fmt.Println("Hasil perbandingan a == b:", a == b)
	fmt.Println("Hasil perbandingan a < b:", a < b)
}