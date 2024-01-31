package main

import (
	"fmt"
)

// declare variable memo using make & map method that have array of ints and will give the return in int data type
var memo = make(map[int]int)

func FibonacciRecursiveDP(n int) int {
	if n <= 1 {
		return n
	}

	// Check if the Fibonacci number is already calculated
	if val, ok := memo[n]; ok {
		return val
	}

	// Calculate the Fibonacci number recursively to combine previous and the current number based on sequence
	f := FibonacciRecursiveDP(n-1) + FibonacciRecursiveDP(n-2)

	// Store the calculated Fibonacci number
	memo[n] = f

	return f
}

func main() {
	var n int
	fmt.Print("Enter the value of N Fibonacci based on sequence: ")
	fmt.Scan(&n)

	fmt.Printf("Fibonacci number at position is %d: %d\n", n, FibonacciRecursiveDP(n))
}
