package main

import (
	"fmt"
	"sort"
	"strconv"
)

func occurredOnce(arr []int) string {
	sort.Ints(arr)

	result := ""

	// Check elemen pertama
	if arr[0] != arr[1] {
		result += strconv.Itoa(arr[0]) + " "
	}

	// Check semua elemen apakah berbeda dari tiap elemen array
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] && arr[i] != arr[i-1] {
			result += strconv.Itoa(arr[i]) + " "
		}
	}

	// Check untuk elemen terakhir
	if arr[len(arr)-2] != arr[len(arr)-1] {
		result += strconv.Itoa(arr[len(arr)-1]) + " "
	}

	return result
}

func main() {
	// limitasi jumlah angka yang dimasukan < 11
	arr := make([]int, 11)
	fmt.Println("Enter 11 numbers:")
	for i := 0; i < 11; i++ {
		fmt.Scan(&arr[i])
	}
	result := occurredOnce(arr)
	fmt.Println(result)
}
