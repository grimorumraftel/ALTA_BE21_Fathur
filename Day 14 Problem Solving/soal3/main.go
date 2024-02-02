package main

import (
	"fmt"
	"sort"
)

func main() {
	var heads, knights int
	for {
		fmt.Scan(&heads, &knights)
		if heads == 0 && knights == 0 {
			break
		}

		h := make([]int, heads)
		k := make([]int, knights)

		for i := 0; i < heads; i++ {
			fmt.Scan(&h[i])
		}

		for i := 0; i < knights; i++ {
			fmt.Scan(&k[i])
		}

		sort.Ints(h)
		sort.Ints(k)

		coins := 0
		possible := true

		for i, j := 0, 0; i < heads; i++ {
			for j < knights && k[j] < h[i] {
				j++
			}

			if j == knights {
				possible = false
				break
			}

			coins += k[j]
			j++
		}

		if possible {
			fmt.Println(coins)
		} else {
			fmt.Println("Loowater is doomed!")
		}
	}
}
