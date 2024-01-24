package main

import "fmt"

func ArrayUnique(arrayA, arrayB []int) []int {
	var muncul = map[int]int{}
	var result = []int{}
	for i := 0; i < len(arrayA); i++ {
		if _, found := muncul[arrayA[i]]; found {
			muncul[arrayA[i]]++
		} else {
			muncul[arrayA[i]] = 0
		}
	}

	for i := 0; i < len(arrayB); i++ {
		if _, found := muncul[arrayB[i]]; found {
			muncul[arrayB[i]]++
		}
	}

	for idx, val := range muncul {
		if val == 0 {
			result = append(result, idx)
		}
	}

	// fmt.Println(muncul)

	return result
}

func main() {
	fmt.Println(ArrayUnique([]int{1, 2, 3, 4}, []int{1, 3, 5, 10, 16}))   // [2 4]
	fmt.Println(ArrayUnique([]int{10, 20, 30, 40}, []int{5, 10, 15, 59})) // [20 30 40]
	fmt.Println(ArrayUnique([]int{1, 3, 7}, []int{1, 3, 5}))              // [7]
	fmt.Println(ArrayUnique([]int{3, 8}, []int{2, 8}))                    // [3]
	fmt.Println(ArrayUnique([]int{1, 2, 3}, []int{3, 2, 1}))              // []
}
