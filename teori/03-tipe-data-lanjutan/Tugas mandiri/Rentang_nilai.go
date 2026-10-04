package main
import ("fmt")

func main() {
	var x int
	var low int
	var high int
	fmt.Scanf("%d", &x)
	fmt.Scanf("%d", &low)
	fmt.Scanf("%d", &high)
	fmt.Println((x >= low) && (x <= high))
}