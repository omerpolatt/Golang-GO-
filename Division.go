package main

import "fmt"

func RockAndRoll(n int) string {
	if n < 0 {
		return "error: number is negative\n"
	}
	if n%3 == 0 && n%2 == 0 {
		return "rock and roll\n"
	}
	if n%3 == 0 {
		return "roll\n"
	}
	if n%2 == 0 {
		return "rock\n"
	}
	return "error: non divisible\n"
}

func main() {
	fmt.Print(RockAndRoll(4))
	fmt.Print(RockAndRoll(9))
	fmt.Print(RockAndRoll(6))
}
