package main
import "fmt"
// kamus: p, l: integer, luas: integer, keliling: integer
// algoritma: input p, l,
// luas ← p * l
// keliling ← 2 * (p + l)
// output luas, keliling
func main() {
	var p, l int
	var luas, keliling int

	fmt.Scan(&p, &l)
	luas = p * l
	keliling = 2 * (p + l)

	fmt.Println(luas)
	fmt.Println(keliling)
}