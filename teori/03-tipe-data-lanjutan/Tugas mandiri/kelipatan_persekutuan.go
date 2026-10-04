package main
import ("fmt")

func main() {
	var (n, a, b int)
	fmt.Scanf("%d %d %d", &n, &a, &b)

	if (n % a == 0 && n % b == 0) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

