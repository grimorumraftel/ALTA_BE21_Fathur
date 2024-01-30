package main

import (
	"fmt"
)

func findMinMax(numbers *[6]int) (int, int) {
	min := (*numbers)[0]
	max := (*numbers)[0]

	for _, num := range *numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}

func main() {
	var numbers [6]int

	fmt.Println("Enter 6 numbers:")

	for i := 0; i < 6; i++ {
		fmt.Scan(&numbers[i])
	}

	min, max := findMinMax(&numbers)

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
