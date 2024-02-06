package main

import (
	"fmt"
	"regexp"
)

func main() {
	var linesCount int
	fmt.Scanln(&linesCount)
	lines := make([]string, linesCount)
	for i := 0; i < linesCount; i++ {
		fmt.Scanln(&lines[i])
	}
	var substringCount int
	fmt.Scanln(&substringCount)
	substrings := make([]string, substringCount)
	for i := 0; i < substringCount; i++ {
		fmt.Scanln(&substrings[i])
	}
	for j := 0; j < len(substrings); j++ {
		count := 0
		pattern := "\\w+" + substrings[j] + "\\w+"
		for i := 0; i < len(lines); i++ {
			line := lines[i]
			re := regexp.MustCompile(pattern)
			matches := re.FindAllString(line, -1)
			count += len(matches)
		}
		fmt.Println(count)
	}
}
