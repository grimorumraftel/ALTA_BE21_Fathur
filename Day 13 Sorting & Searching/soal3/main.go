package main

import (
	"fmt"
	"math"
)

func playingDomino(dominoCards [][]int, deckCards [][]int) []int {
	length := len(dominoCards)
	left := make([]int, length)
	right := make([]int, length)
	res := []int{}

	for i := length - 1; i >= 0; i-- {
		if dominoCards[i][0] == 0 {
			left[i] = 0
		} else if dominoCards[i][0] == -1 {
			if i != length-1 && left[i+1] != math.MaxInt32 {
				left[i] = left[i+1] + 1
			}
		}

		if dominoCards[i][1] == 1 {
			right[i] = 0
		} else if dominoCards[i][1] == -1 {
			if i != 0 && right[i-1] != math.MaxInt32 {
				right[i] = right[i-1] + 1
			}
		}
	}

	for i := 0; i < length; i++ {
		if right[i] < left[i] {
			res = append(res, 1)
		} else if right[i] > left[i] {
			res = append(res, 0)
		} else {
			res = append(res, -1)
		}
	}

	return res
}

func main() {
	dominoCards := [][]int{{1, 2}, {3, 4}, {5, 6}}
	deckCards := [][]int{{7, 8}, {9, 10}, {11, 12}}
	result := playingDomino(dominoCards, deckCards)
	fmt.Println(result)
}
