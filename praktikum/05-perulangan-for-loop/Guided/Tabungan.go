package main
import "fmt"
func main() {
	var a, b, n int
	fmt.Scan(&a, &b, &n)
	
	total := 0;
	tabunganHarian := a;

	for i := 1; i <= n; i++ {
		total += tabunganHarian
		tabunganHarian += b
	}
	fmt.Println(total)
}
