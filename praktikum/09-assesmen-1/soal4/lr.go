package main
import "fmt"

func main() {
	var L, R int
	var jumlah int
	fmt.Scan(&L, &R)

	for i := L; i <= R; i++ {
		angka := i
		jumlahDigit := 0
		sebelumnya := 10
		valid := true

		for angka > 0 {
			digit := angka % 10

			jumlahDigit += digit

			if digit > sebelumnya {
				valid = false
			}
			sebelumnya = digit
			angka /= 10
		}
		prima := true

		if jumlahDigit < 2 {
			prima = false
		} else {
			for j := 2; j < jumlahDigit; j++ {
				if jumlahDigit%j == 0{
				prima = false
			}
		}
	}
	if valid && prima {
		jumlah++
	}
	
}
fmt.Println(jumlah)
}