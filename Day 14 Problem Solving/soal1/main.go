package main

import (
	"fmt"
)

func simpleEquations() {
	var A, B, C int

	fmt.Print("Enter the value for A: ")
	fmt.Scan(&A)

	fmt.Print("Enter the value for B: ")
	fmt.Scan(&B)

	fmt.Print("Enter the value for C: ")
	fmt.Scan(&C)

	for x := 1; x <= A; x++ {
		for y := 1; y <= B; y++ {
			for z := 1; z <= C; z++ {
				if x+y+z == A && x*y*z == B && x*x+y*y+z*z == C {
					fmt.Printf("x = %d, y = %d, z = %d\n", x, y, z)
					return
				}
			}
		}
	}

	fmt.Println("No solution found")
}

func main() {
	simpleEquations()
}
