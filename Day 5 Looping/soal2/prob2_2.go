package main

import "fmt"

func main() {
	var n int
	fmt.Println("Masukan angka yang ingin diketahui faktor bilangannya: ")
	fmt.Scanln(&n)
	for h := n; h > 0; h-- {
		if n%h == 0 {
			fmt.Println(h)
		}
	}
}
