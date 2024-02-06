package main

import (
	"fmt"
)

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	var T, n int
	fmt.Scan(&T)
	for T > 0 {
		T--
		fmt.Scan(&n)
		t := make([]int, n)
		best := make([]int64, n+3)
		suff := make([]int64, n+3)
		for i := n; i < n+3; i++ {
			best[i] = 0
			suff[i] = 0
		}
		for i := 0; i < n; i++ {
			fmt.Scan(&t[i])
		}
		for i := n - 1; i >= 0; i-- {
			suff[i] = suff[i+1] + int64(t[i])
		}
		for i := n - 1; i >= 0; i-- {
			best[i] = suff[i] - min(best[i+1], min(best[i+2], best[i+3]))
		}
		fmt.Println(best[0])
	}
}
