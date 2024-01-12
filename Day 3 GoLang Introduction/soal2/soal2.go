// Problem 4 (Menghitung Luas Segitiga)
package main

import "fmt"

func main() {
	var alas, tinggi, hasil float64
	fmt.Print("Silahkan masukan panjang alas : ")
	fmt.Scan(&alas)
	fmt.Print("Silahkan masukan panjang tinggi : ")
	fmt.Scan(&tinggi)
	hasil = alas * tinggi / 2
	fmt.Printf("Luas segitiga adalah %f", hasil)

}
