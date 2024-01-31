package main

import (
	"fmt"
)

func findMaxSum(arr []int) int {
	maxSum := 0
	currentSum := 0

	for _, num := range arr {
		currentSum += num
		if currentSum < 0 {
			currentSum = 0
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

func main() {
	var n int
	fmt.Print("Enter the number of elements in the array: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Enter the elements of the array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	maxSum := findMaxSum(arr)
	fmt.Println("Maximum possible sum:", maxSum)
}
