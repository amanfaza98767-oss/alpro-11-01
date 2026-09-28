package main
import "fmt"
// kamus: c: real, f: real
// algoritma: input c,
// f ← (c * 9 / 5) + 32
// output f
func main() {
	var c float64
	var f float64

	fmt.Scan(&c)

	f = (c * 9 / 5) + 32

	fmt.Println(f)
}