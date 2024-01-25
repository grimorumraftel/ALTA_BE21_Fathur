package main

import (
	"fmt"
	"strings"
)

type SubstitutionCipher struct {
	shift int
}

func (sc *SubstitutionCipher) Encrypt(text string) string {
	runes := []rune(text)

	offset := func(char rune) rune {
		if char >= 'a' && char <= 'z' {
			return 'a' + ((char - 'a' + rune(sc.shift)) % 26)
		} else if char >= 'A' && char <= 'Z' {
			return 'A' + ((char - 'A' + rune(sc.shift)) % 26)
		}
		return char
	}

	encoded := strings.Map(offset, string(runes))
	return strings.ToLower(encoded)
}

func (sc *SubstitutionCipher) Decrypt(text string) string {
	runes := []rune(text)

	offset := func(char rune) rune {
		if char >= 'a' && char <= 'z' {
			return 'a' + ((char - 'a' - rune(sc.shift) + 26) % 26)
		} else if char >= 'A' && char <= 'Z' {
			return 'A' + ((char - 'A' - rune(sc.shift) + 26) % 26)
		}
		return char
	}

	decoded := strings.Map(offset, string(runes))
	return strings.ToLower(decoded)
}

func main() {
	sc := SubstitutionCipher{}

	var choice int
	var input string

	for {
		fmt.Println("1. Encrypt")
		fmt.Println("2. Decrypt")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter a string to encrypt: ")
			fmt.Scanln(&input)
			fmt.Print("Enter the number of positions to move forwards the string: ")
			fmt.Scanln(&sc.shift)
			encrypted := sc.Encrypt(input)
			fmt.Println("Encrypted text:", encrypted)
		case 2:
			fmt.Print("Enter a string to decrypt: ")
			fmt.Scanln(&input)
			fmt.Print("Enter the number of positions to move backwards the string: ")
			fmt.Scanln(&sc.shift)
			decrypted := sc.Decrypt(input)
			fmt.Println("Decrypted text:", decrypted)
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please only choose the number on the screen!.")
		}

		fmt.Println()
	}
}
