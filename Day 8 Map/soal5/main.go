package main

import (
	"fmt"
)

func removeDuplicates(arr []int) []int {
	mapVar := map[int]bool{}
	result := []int{}
	for _, e := range arr {
		if !mapVar[e] {
			mapVar[e] = true
			result = append(result, e)
		}
	}
	return result
}

func main() {
	var n int
	fmt.Println("Enter lenght of array: ")
	fmt.Scanln(&n)

	arr := make([]int, n)
	fmt.Println("Enter the elements of the array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	result := removeDuplicates(arr)
	fmt.Println("The new array result after remove the duplicate is:", result)
}
