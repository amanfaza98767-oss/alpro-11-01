package main
	import (
		"fmt"
	)
	func main() {
	var angka int
	
	fmt.Print("masukan bilangan poitif: ")
	fmt.Scan(&angka)

	ascending := true
	descending := true 
	pertama := true
	
	var digitSebelumnya  int 

	for angka > 0 {
		digit := angka % 10

		if !pertama {
			if digit < digitSebelumnya {
				descending = false
			}
			if digit == digitSebelumnya{
				descending = false
				ascending  = false
			}
			if digit > digitSebelumnya {
				ascending = false
			}
		}
		digitSebelumnya = digit
		pertama = false
		angka = angka / 10
	}
	if ascending {
		fmt.Println("ascending")
	} else if  descending {
		fmt.Println("descending")
	} else {
		fmt.Println("tidak terurut")
	}

	}