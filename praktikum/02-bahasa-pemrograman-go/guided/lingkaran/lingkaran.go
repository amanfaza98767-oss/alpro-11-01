package main
import "fmt"

func main() {
	var pi float64 = 3.14
	var jariJari float64

	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&jariJari)

	var luas float64 = pi * jariJari * jariJari
	var keliling float64 = 2 * pi * jariJari

	fmt.Printf("Luas lingkaran: %.2f\n", luas)
	fmt.Printf("Keliling lingkaran: %.2f\n", keliling)
}