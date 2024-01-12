// Problem 6 (Program Menghitung Harga Setelah Diskon)
package main

import "fmt"

func main() {
	var hargaBarang, totalBayar float64
	fmt.Print("Silahkan masukan Harga Barang : ")
	fmt.Scan(&hargaBarang)

	if hargaBarang >= 200000 && hargaBarang <= 350000 {
		totalBayar = hargaBarang - (0.1 * hargaBarang)
	} else if hargaBarang == 370000 {
		totalBayar = hargaBarang - (0.15 * hargaBarang)
	} else if hargaBarang >= 400000 && hargaBarang <= 550000 {
		totalBayar = hargaBarang - (0.2 * hargaBarang)
	} else if hargaBarang >= 600000 && hargaBarang <= 750000 {
		totalBayar = hargaBarang - (0.3 * hargaBarang)
	} else if hargaBarang >= 800000 {
		totalBayar = hargaBarang - (0.4 * hargaBarang)
	} else {
		totalBayar = hargaBarang
	}

	fmt.Printf("Total harga yang harus dibayar adalah Rp. %f", totalBayar)
}
