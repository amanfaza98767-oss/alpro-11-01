package main
import ("fmt")

func main() {
	var warna string
	fmt.Print("Masukkan warna lampu lalu lintas (merah, kuning, hijau): ")
	fmt.Scan(&warna)

	switch warna {
	case "merah":
		fmt.Println("Lampu merah: Berhenti")
	case "kuning":
		fmt.Println("Lampu kuning: Hati-hati")
	case "hijau":
		fmt.Println("Lampu hijau: Jalan")
	default:
		fmt.Println("Warna lampu tidak valid")
	}
}