package main

import "fmt"

func main() {
	var number int = 0
	var inputan int = 0

	fmt.Print("Enter Number: ")
	fmt.Scanf("%d", &number)

	for count := 2; count < number/2; count++ {
		if number%count == 0 {
			inputan = 1

		}
	}
	if inputan == 1 {
		fmt.Print(number, " Bukanlah Bilangan Prima")
	} else {
		fmt.Print(number, " Adalah Bilangan Prima")
	}
}
