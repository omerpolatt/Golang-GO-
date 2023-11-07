package main

import (
	"fmt"
	"os"
	"strconv"
)

func ispowerof2(n int) bool {
	if n <= 0 {
		return false
	}

	for i := 1; i <= n; i *= 2 {
		if i == n {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args

	if len(args) == 2 {
		input, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println()

		}

		cikti := ispowerof2(input)

		fmt.Println(cikti)
	}
}
