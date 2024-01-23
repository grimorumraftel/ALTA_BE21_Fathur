package main

import (
	"fmt"
	"math"
)

func exponential(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}

func main() {
	var base, exponent float64
	fmt.Print("Masukan angka: ")
	fmt.Scan(&base)
	fmt.Print("Masukan jumlah pangkat exponen: ")
	fmt.Scan(&exponent)
	result := exponential(base, exponent)
	fmt.Println("Result:", result)
}
