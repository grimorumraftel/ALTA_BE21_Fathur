package main

import "fmt"

func binarySearch(arr []int, search int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == search {
			return mid
		} else if arr[mid] < search {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	var n int
	fmt.Print("Enter the number of array elements: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Enter the elements of array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	var search int
	fmt.Print("Enter the value in array list to search: ")
	fmt.Scan(&search)

	index := binarySearch(arr, search)
	if index != -1 {
		fmt.Printf("Value %d found at index %d\n", search, index)
	} else {
		fmt.Printf("Value %d not found in the array\n", search)
	}
}
