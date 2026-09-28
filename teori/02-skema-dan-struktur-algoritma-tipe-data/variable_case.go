package main
import ("fmt")

func main() {
	var name string

	name = "Amanullah Risky Maya Sekti"

	fmt.Println("Nama : ", name)

	var lastName string = "Maya Sekti"
	fmt.Println("Nama Belakang : ", lastName)

	middleName := "Risky"
	fmt.Println("Nama Tengah : ", middleName)

	var(
		fullName = "Amanullah Risky Maya Sekti"
		firstName = "Amanullah"
	)

	fmt.Println(fullName)
	fmt.Println(firstName)
}