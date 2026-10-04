package main
import ("fmt")

func main() {
	var x int
	var low int
	var high int
	fmt.Scanf("%d", &x)
	fmt.Scanf("%d", &low)
	fmt.Scanf("%d", &high)
	if (x >= low) && (x <= high) {
		fmt.Println("True")
	}else {
		fmt.Println("False")
	}
}