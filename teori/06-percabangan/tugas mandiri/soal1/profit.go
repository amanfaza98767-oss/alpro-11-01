package main
import "fmt"

func main() {
	var sebelumnya, sekarang float64
	fmt.Scan(&sebelumnya, &sekarang)

	if sekarang > sebelumnya {
		fmt.Printf("peningkatan sebesar %.0f\n", sekarang-sebelumnya)
	} else if sekarang < sebelumnya {
		fmt.Printf("penurunan sebesar %.0f\n", sebelumnya-sekarang) 
	} else {
		fmt.Println("tetap")
	}
}