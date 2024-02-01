package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func countMostFrequent(arr []int) string {
	frequency := make(map[int]int)
	for _, num := range arr {
		frequency[num]++
	}

	type elementFreq struct {
		element int
		freq    int
	}

	var elements []elementFreq
	for num, freq := range frequency {
		elements = append(elements, elementFreq{num, freq})
	}

	sort.Slice(elements, func(i, j int) bool {
		return elements[i].freq > elements[j].freq
	})

	var mostFrequentElements []string
	for _, ef := range elements {
		mostFrequentElements = append(mostFrequentElements, strconv.Itoa(ef.element))
	}

	return strings.Join(mostFrequentElements, ", ")
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

	mostFrequentElements := countMostFrequent(arr)
	fmt.Println("Most frequently appearing elements (sorted by frequency):", mostFrequentElements)

	// Display the count for each element
	frequency := make(map[int]int)
	for _, num := range arr {
		frequency[num]++
	}

	fmt.Println("Frequency of elements:")
	for _, ef := range strings.Split(mostFrequentElements, ", ") {
		element, _ := strconv.Atoi(ef)
		fmt.Printf("%d: %d\n", element, frequency[element])
	}
}
