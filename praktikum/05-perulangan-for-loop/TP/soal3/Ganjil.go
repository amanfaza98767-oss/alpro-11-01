package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	jumlah := 0
	for i := 1; i <= n; i += 2 {
		fmt.Println(i, "")
		jumlah += i
	}
	fmt.Println("Jumlah bilangan ganjil dari 1 sampai", n, "adalah", jumlah)
}