package main

import (
	"fmt"
	"os"
)

func repeatAlpha(s string) string {
	result := ""
	for _, char := range s {
		if isAlphabetic(char) {
			repeatCount := int(char - 'a' + 1)
			for i := 0; i < repeatCount; i++ {
				result += string(char)
			}
		}

	}
	return result
}

func isAlphabetic(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
	} else {
		input := os.Args[1]
		result := repeatAlpha(input)
		fmt.Println(result)
	}
}
