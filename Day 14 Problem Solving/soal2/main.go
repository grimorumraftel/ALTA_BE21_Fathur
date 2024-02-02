package main

import (
	"fmt"
	"sort"
)

type Denomination struct {
	Value int
	Count int
}

func moneyChange(amount int, denominations []int) []Denomination {
	sort.Slice(denominations, func(i, j int) bool {
		return denominations[i] > denominations[j]
	})

	result := []Denomination{}
	for _, denom := range denominations {
		count := amount / denom
		if count > 0 {
			result = append(result, Denomination{
				Value: denom,
				Count: count,
			})
			amount %= denom
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})

	return result
}

func main() {
	var amount int
	fmt.Print("Enter the amount of money: ")
	fmt.Scan(&amount)

	denominations := []int{1, 20, 50, 100, 200, 500, 1000, 2000, 5000, 1000, 2000, 5000, 10000, 20000, 50000, 100000}

	result := moneyChange(amount, denominations)

	fmt.Println("Denominations:")
	denominationList := []int{}
	for _, denom := range result {
		for i := 0; i < denom.Count; i++ {
			denominationList = append(denominationList, denom.Value)
		}
	}
	fmt.Println(denominationList)
}
