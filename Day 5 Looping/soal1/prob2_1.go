package main

import "fmt"

func main() {
	var number int
	fmt.Println("Masukan angka yang ingin diketahui faktor bilangannya: ")
	fmt.Scanln(&number)
	for n := 1; n <= number; n++ {
		if number%n == 0 {
			fmt.Println(n)
		}

	}
}
