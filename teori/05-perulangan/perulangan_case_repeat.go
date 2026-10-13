package main
import "fmt"

func main() {
	var n, i int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)
	i = 1
	for {
		fmt.Println(i)
		i++
		if i > n {
			break
		}
	}
}