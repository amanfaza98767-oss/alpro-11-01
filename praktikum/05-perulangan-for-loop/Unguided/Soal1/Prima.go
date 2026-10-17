package main
import "fmt"
func main () {
	var n int
	fmt.Scan(&n)

	var jumlahFaktor int

	for i :=1; i <=n; i++{
	if n%i == 0 {
		fmt.Println(i)
		jumlahFaktor++
	}
}

if jumlahFaktor == 2 {
	fmt.Println("ya")

} else {
	fmt.Println("tidak")

}
}