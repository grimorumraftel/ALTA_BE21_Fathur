package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	var a, b int

	fmt.Print("Enter 1st number: ")
	fmt.Scan(&a)

	fmt.Print("Enter 2nd number: ")
	fmt.Scan(&b)

	swap(&a, &b)

	fmt.Println("After swapping:")
	fmt.Println("1st number:", a)
	fmt.Println("2nd number:", b)
}
