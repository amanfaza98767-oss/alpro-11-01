package main
import "fmt"
// kamus: x:real, f:real
// algoritma: input x, 
// f ← 3*x*x + 5*x - 2
// output f
func main() {
	var x float64
	var f float64

	fmt.Scan(&x)
	//menghitung f(x) = 3x² + 5x − 2
	f = (x*x+2*x+1) / (x-3)

	fmt.Println(f)
}