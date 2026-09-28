package main

func main() {
	var umur int8
	var suhu float64

	suhu = 36.3
	umur = 10

	println("Umur : ", umur)
	println("Suhu : ", suhu)
	println("Alamat memori dari var suhu : ", &suhu)
	println("Alamat memori dari var umur : ", &umur)
}
