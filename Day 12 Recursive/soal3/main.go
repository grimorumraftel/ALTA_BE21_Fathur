package main

import (
	"fmt"
	"math"
	"strconv"
)

func isPrime(n int64) bool {
	if n <= 1 {
		return false
	}

	for i := int64(2); i <= int64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var height, width int
	var start int64
	fmt.Print("Enter the height of the square: ")
	fmt.Scan(&height)
	fmt.Print("Enter the width of the square: ")
	fmt.Scan(&width)
	fmt.Print("Enter the starting prime number to print: ")
	fmt.Scan(&start)

	var primeSum int64
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			num := start + int64(i*width+j)
			if isPrime(num) {
				primeSum += num
			}
		}
	}
	fmt.Println("Sum of prime numbers:", strconv.FormatInt(primeSum, 10))
}
