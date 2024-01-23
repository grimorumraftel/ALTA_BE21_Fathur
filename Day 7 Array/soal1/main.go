package main

import "fmt"

func isThisPrima(angka int) {
	pembagi := 0
	for i := 1; i <= angka; i++ {
		if angka%i == 0 {
			pembagi++
		}
	}
	if pembagi == 2 {
		fmt.Println("prima")
	} else {
		fmt.Println("bukan prima")
	}
}

func main() {
	fmt.Print("Masukan angka yang ingin di cek bilangan prima atau bukan: ")
	var x int
	fmt.Scan(&x)
	isThisPrima(x)
}
