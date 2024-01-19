package main

import "fmt"

func tabelPerkalian(n int) {
	if n < 1 || n > 30 {
		fmt.Print("Inputan Angka Salah")
		return
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Printf("%6d", i*j)
		}
		fmt.Println()
	}
}

func main() {
	fmt.Print("Masukan jumlah baris angka perkalian: ")
	var x int
	fmt.Scan(&x)
	tabelPerkalian(x)
}
