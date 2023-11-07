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
	lastWord := ""

	for i := len(input) - 1; i >= 0; i-- {
		if input[i] != ' ' {
			lastWord = string(input[i]) + lastWord
		} else if lastWord != "" {
			break
		}
	}

	if lastWord == "" {
		fmt.Println()
		return
	}

	fmt.Println(lastWord)
}
