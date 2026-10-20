package main
import "fmt"
func main() {
 var nilai int

 for {
	fmt.Print("masukan nilai anda: ")
	fmt.Scan(&nilai)
	if nilai == -1 {
		fmt.Println("program selesai")
		break
	}
	if nilai >= 60 {
		fmt.Println("anda lulus ujian")
	} else {
		fmt.Println("anda tidak lulus ujian")
	}
 }
}