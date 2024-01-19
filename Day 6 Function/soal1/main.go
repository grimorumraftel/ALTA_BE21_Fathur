// PROB 3-1
package main

import "fmt"

func asteriskPyramid(n int) {
	for x := 1; x <= n; x++ {
		for y := n; y > x; y-- {
			fmt.Print(" ")
		}
		for z := 1; z <= x; z++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("Masukan Tinggi Segitiga: ")
	var n int
	fmt.Scanln(&n)
	asteriskPyramid(n)
}
