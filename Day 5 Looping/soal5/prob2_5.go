package main

import "fmt"

func pow(angka, pangkat int) int {
	result := 1
	for i := 0; i < pangkat; i++ {
		result *= angka
	}
	return result
}
func main() {
	var angka, pangkat, hasilakhir int
	fmt.Print("Masukan Angka: ")
	fmt.Scanln(&angka)
	fmt.Print("Masukan Pangkat: ")
	fmt.Scanln(&pangkat)
	hasilakhir = pow(angka, pangkat)
	fmt.Print("Hasil pangkat exponential adalah ", hasilakhir)
}
