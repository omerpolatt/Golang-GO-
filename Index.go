package main

import "fmt"

func Index(s string, toFind string) int {
	// Our function that searches for the desired fragment in the given String
	a := len(s)
	b := len(toFind)
	if a < b {
		return -1
	}
	for i := 0; i <= a-b; i++ {
		if s[i:i+b] == toFind {
			return i
		}
	}
	return -1
}

func main() {
	output := Index("Hello", "lo")
	fmt.Print(output)
}
