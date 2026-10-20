package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	var biaya float64 = 10000000
	i := 1

	for i < n {
		biaya = biaya * 1.06
		i++
	}

	fmt.Printf("%.0f\n", biaya)
}