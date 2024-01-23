package main

import (
	"fmt"
)

func removeDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func main() {
	var input1, input2 []int

	fmt.Println("Enter the first array:")
	for {
		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}
		input1 = append(input1, num)
	}

	fmt.Println("Enter the second array:")
	for {
		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}
		input2 = append(input2, num)
	}

	// Join kedua array
	joinedSlice := append(input1, input2...)

	// Remove duplicate pada kedua array
	uniqueSlice := removeDuplicate(joinedSlice)

	// print slice baru dari kedua array yg sudah di join
	fmt.Println("Joined and unique slice:", uniqueSlice)
}
