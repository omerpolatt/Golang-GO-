package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	input := os.Args[1]
	firstWord := ""

	for i := 0; i < len(input); i++ {
		char := rune(input[i])

		if char != ' ' {
			firstWord += string(char)
		}
		if firstWord != "" {
			break
		}
	}

	if firstWord == "" {
		fmt.Println()
		return
	}

	fmt.Println(firstWord)
}
