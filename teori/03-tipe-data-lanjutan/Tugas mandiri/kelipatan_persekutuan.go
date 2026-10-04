package main
import ("fmt")

func main() {
	var (n, a, b int)
	fmt.Scanf("%d %d %d", &n, &a, &b)
fmt.Println((n%a == 0) && (n%b == 0))
}

