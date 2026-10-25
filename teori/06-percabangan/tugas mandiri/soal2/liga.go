package main
import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	terbanyak := a
	tersedikit := a

	if b > terbanyak {
		terbanyak = b
	}
	if c > terbanyak {
		terbanyak = c
	}
	if d > terbanyak {
		terbanyak = d
	}
	if b < tersedikit {
		tersedikit = b
	}
	if c < tersedikit {
		tersedikit = c
	}
	if d < tersedikit {
		tersedikit = d
	}

	fmt.Println(terbanyak, tersedikit)
}