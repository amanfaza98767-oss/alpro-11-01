package main
import "fmt"

func main() {
var totalHari int

fmt.Scan(&totalHari)

tahun := totalHari / 360
sisa := totalHari % 360

bulan := sisa / 30
sisa = sisa % 30

minggu := sisa / 7
sisa = sisa % 7

fmt.Printf("%d %d %d %d\n", tahun, bulan, minggu, sisa)
}
