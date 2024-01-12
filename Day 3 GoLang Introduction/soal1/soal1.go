// Problem 3  (Program Hello World + String Nama)
package main

import "fmt"

func main() {
	var nama string
	fmt.Print("Silahkan Tuliskan Nama Anda: ")
	fmt.Scan(&nama)
	fmt.Printf("Hello, %s! Saya Golang, bahasa yang sangat menyenangkan.\n", nama)
}
