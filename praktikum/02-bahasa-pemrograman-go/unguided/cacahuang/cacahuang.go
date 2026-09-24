package main
import ("fmt")

func main() {
	var totalUang int

	//membaca masukan nominal uang
	fmt.Scan(&totalUang)

	//menghitung lembar puluhan ribu
	sepuluhribu := totalUang / 10000
	sisa := totalUang % 10000

	//menghitung lembar lima ribu
	limaribu := sisa / 5000
	sisa = sisa % 5000

	//menghitung lembar  seriburibu
	seribu := sisa / 1000
	
	//menampilkan hasil keluaran sesuai format
	fmt.Printf("%d %d %d\n", sepuluhribu, limaribu, seribu)
}

