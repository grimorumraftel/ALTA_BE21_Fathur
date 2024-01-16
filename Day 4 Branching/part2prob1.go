package main

import "fmt"

func main() {
	var score int
	fmt.Println("Masukan Score Akhir : ")
	fmt.Scan(&score)

	if score >= 80 && score <= 100 {
		fmt.Print("Nilai A")
	} else if score >= 65 && score <= 79 {
		fmt.Print("Nilai B+")
	} else if score >= 50 && score <= 64 {
		fmt.Print("Nilai B")
	} else if score >= 35 && score <= 49 {
		fmt.Print("Nilai C")
	} else if score >= 0 && score <= 34 {
		fmt.Print("Nilai D")
	} else {
		fmt.Print("Angka yang anda masukan salah!")
	}
}
