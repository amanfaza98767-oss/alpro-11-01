package main
import "fmt"
func main () {
	var x, digit, kuadrat, total int
	fmt.Scan(&x)
	
	total = 0

	for x > 0 {
		digit = x % 10
		kuadrat = digit * digit

		fmt.Println(kuadrat)
		total += kuadrat
		x = x/10
	}
	fmt.Println(total)
}