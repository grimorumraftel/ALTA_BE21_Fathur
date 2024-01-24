package main

import (
	"fmt"
)

// function maxSubArrayOfSizeK
func maxSubArrayOfSizeK(arr []int, windowSize int) (int, []int) {
	var maximumSum, windowSum, start, end int
	var subArray []int

	for end = 0; end < len(arr); end++ {
		windowSum += arr[end]
		if end >= windowSize-1 {
			if maximumSum <= windowSum {
				maximumSum = windowSum
				subArray = arr[start : end+1]
			}
			windowSum -= arr[start]
			start++
		}
	}

	return maximumSum, subArray
}

func main() {
	var size, k int
	fmt.Print("Enter the size of the array: ")
	fmt.Scan(&size)

	arr := make([]int, size)
	fmt.Println("Enter the elements of the array:")
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Print("Enter the window size (k): ")
	fmt.Scan(&k)

	maxSum, subArray := maxSubArrayOfSizeK(arr, k)
	fmt.Println("Maximum sum of subarray of size", k, "is", maxSum)
	fmt.Println("Array elements that make up the result:", subArray)
}
