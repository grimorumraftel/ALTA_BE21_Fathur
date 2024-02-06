package main

import (
	"fmt"
)

func repeatedDigit(n int) int {
	s := make(map[int]bool)

	for n != 0 {
		d := n % 10
		if s[d] {
			return 0
		}
		s[d] = true
		n = n / 10
	}
	return 1
}

func calculate(L, R int) int {
	answer := 0
	for i := L; i < R+1; i++ {
		answer = answer + repeatedDigit(i)
	}

	return answer
}

func main() {
	L := 10
	R := 20
	result := calculate(L, R)
	fmt.Println(result)
}
