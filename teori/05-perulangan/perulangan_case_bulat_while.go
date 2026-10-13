package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for iterasi := 1; iterasi <= n; iterasi++ {
		fmt.Println(iterasi)
	}
}