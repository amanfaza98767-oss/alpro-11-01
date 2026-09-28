package main

import "fmt"
// kamus: x, y : integer
// algoritma: input x, y, 
// f ← 5*x*x - 2*x*y + y*y*y / (x + 1)
// output f
func main() {
	var x, y int
	var f float64

	fmt.Scan(&x, &y)
//menghitung f(x, y) = 5x² − 2xy + y³ / (x + 1) 
	f = 5*float64(x)*float64(x) - 2*float64(x)*float64(y) + float64(y*y*y)/float64(x+1)
	fmt.Println(f)

}
