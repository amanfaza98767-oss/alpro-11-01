package main
import "fmt"
func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	for i := 1; i <= int(b); i++ {
		fmt.Println(a, "x", i, "=", a*float64(i))
	} 
	fmt.Println("Perkalian", a, "sampai", b, "adalah", a*b)
}