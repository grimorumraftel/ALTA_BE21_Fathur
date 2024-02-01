package main

import (
	"fmt"
)

func findMaxMin(arr []int) (int, int, int, int) {
	max := arr[0]
	min := arr[0]
	maxIndex := 0
	minIndex := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			maxIndex = i
		}
		if arr[i] < min {
			min = arr[i]
			minIndex = i
		}
	}

	return max, maxIndex, min, minIndex
}

func main() {
	var n int
	fmt.Print("Enter the number of elements in the array: ")
	fmt.Scan(&n)

	array := make([]int, n)
	fmt.Println("Enter the elements of the array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	maximum, maxIndex, minimum, minIndex := findMaxMin(array)
	fmt.Println("Minimum:", minimum, "at index", minIndex)
	fmt.Println("Maximum:", maximum, "at index", maxIndex)
}
