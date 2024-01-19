package main

import "fmt"

func main() {
	var input uint8
	fmt.Scanln(&input)
	var isPrima bool = true
	for i := 2; i < int(input); i++ {
		if int (input)%i == 0 {
			isPrima = false
		}
	}

	if isPrima {
		fmt.Println("Prima")
	}
}
