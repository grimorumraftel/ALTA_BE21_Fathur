// Problem 5 (Menghitung Luas Permukaan Tabung)
package main

import "fmt"

func main() {
	const Phi = 3.14
	var T, r, hasil float64
	fmt.Print("Silahkan masukan tinggi tabung : ")
	fmt.Scan(&T)
	fmt.Print("Silahkan masukan jari-jari : ")
	fmt.Scan(&r)
	hasil = (2 * Phi * r * r) + (2 * Phi * r * T)
	fmt.Printf("Luas permukaan tabung adalah %f", hasil)
}
