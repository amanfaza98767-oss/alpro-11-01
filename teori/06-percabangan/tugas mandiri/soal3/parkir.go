package main
import "fmt"

func main() {
	var h1, m1, h2, m2 int
	fmt.Scan(&h1, &m1, &h2, &m2)

	if h1 < 7 {
		h1 += 12
	}
	if h2 < 7 {
		h2 += 12
	}

	total1 := h1*60 + m1
	total2 := h2*60 + m2

	durasi := total2 - total1
	jam := durasi / 60
	menit := durasi % 60

	fmt.Println(jam, "jam", menit, "menit")
}