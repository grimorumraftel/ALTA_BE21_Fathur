package main

import (
	"fmt"
	"sort"
)

func countMaxItems(cost []float64, budget float64) int {
	sort.Float64s(cost)

	pre := make([]float64, len(cost))
	ans := 0

	for i := 0; i < len(cost); i++ {
		if i > 0 {
			pre[i] += pre[i-1] + cost[i]
		} else {
			pre[i] = cost[i]
		}

		if pre[i] <= budget {
			ans = i + 1
		}
	}

	return ans
}

func main() {
	var budget float64

	fmt.Print("Enter the budget: ")
	fmt.Scan(&budget)

	var limit int
	fmt.Print("Enter the limit: ")
	fmt.Scan(&limit)

	var cost []float64
	fmt.Print("Enter the costs (separated by spaces): ")
	fmt.Scan(&cost)

	for i := 0; i < limit; i++ {
		var c float64
		_, err := fmt.Scan(&c)
		if err != nil {
			break
		}
		cost = append(cost, c)
	}

	maxItems := countMaxItems(cost, budget)
	fmt.Println("Maximum items that can be bought:", maxItems)
}
