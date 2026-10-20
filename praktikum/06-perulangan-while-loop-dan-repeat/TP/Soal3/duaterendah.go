package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var nama string
	var nilai int

	var namaTerendah, namaKedua string
	var nilaiTerendah, nilaiKedua int

	i := 1
	for i <= n {
		fmt.Scan(&nama, &nilai)

		if i == 1 {
			namaTerendah = nama
			nilaiTerendah = nilai
		} else if i == 2 {
			if nilai < nilaiTerendah {
				namaKedua = namaTerendah
				nilaiKedua = nilaiTerendah

				namaTerendah = nama
				nilaiTerendah = nilai
			} else {
				namaKedua = nama
				nilaiKedua = nilai
			}
		} else {
			if nilai < nilaiTerendah {
				namaKedua = namaTerendah
				nilaiKedua = nilaiTerendah

				namaTerendah = nama
				nilaiTerendah = nilai
			} else if nilai < nilaiKedua {
				namaKedua = nama
				nilaiKedua = nilai
			}
		}

		i++
	}

	fmt.Println(namaTerendah)
	fmt.Println(namaKedua)
}