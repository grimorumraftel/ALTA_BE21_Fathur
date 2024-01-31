package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primeX(x int) int {
	count := 0
	num := 2
	for count < x {
		if isPrime(num) {
			count++
			if count == x {
				return num
			}
		}
		num++
	}
	return -1
}

func main() {
	var input int
	fmt.Print("Enter a number: ")
	fmt.Scan(&input)

	fmt.Println("Prime number:")
	primeNumber := primeX(input)
	fmt.Printf("%d\n", primeNumber)
}
