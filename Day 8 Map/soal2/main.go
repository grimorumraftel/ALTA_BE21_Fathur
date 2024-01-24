package main

import (
	"fmt"
	"strings"
)

func cipher(text string, shift int) string {
	runes := []rune(text)

	offset := func(char rune) rune {
		if char >= 'a' && char <= 'z' {
			return 'a' + ((char - 'a' + rune(shift)) % 26)
		} else if char >= 'A' && char <= 'Z' {
			return 'A' + ((char - 'A' + rune(shift)) % 26)
		}
		return char
	}

	encoded := strings.Map(offset, string(runes))
	return strings.ToLower(encoded)
}

func main() {
	fmt.Print("Enter a string: ")
	var input string
	fmt.Scanln(&input)

	fmt.Print("Enter the number offset of string that you want to move forwards: ")
	var offset int
	fmt.Scanln(&offset)

	encoded := cipher(input, offset)
	fmt.Println(encoded)
}
