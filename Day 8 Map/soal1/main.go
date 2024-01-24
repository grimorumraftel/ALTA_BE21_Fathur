package main

import (
	"fmt"
	"strings"
)

func findCommonSubstring(s1 string, s2 string) string {
	m := make(map[rune]bool)

	// store each character of s1 in the map
	for _, char := range s1 {
		m[char] = true
	}

	// Check for common characters in s2
	var commonSubstring strings.Builder
	for _, char := range s2 {
		if m[char] {
			commonSubstring.WriteRune(char)
		}
	}

	return commonSubstring.String()
}

func main() {
	var str1, str2 string
	fmt.Print("Enter the first string: ")
	fmt.Scanln(&str1)
	fmt.Print("Enter the second string: ")
	fmt.Scanln(&str2)

	common := findCommonSubstring(str1, str2)
	if common != "" {
		// Convert the common substring to lowercase
		common = strings.ToLower(common)

		// Check if both strings have the same first alphabet
		if strings.ToLower(string(str1[0])) != strings.ToLower(string(str2[0])) {
			fmt.Println("The first alphabet between two strings must be same")
		} else {
			fmt.Println(common)
		}
	} else {
		fmt.Println("No common substring found between the two strings.")
	}
}
