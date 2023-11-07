package main

import "fmt"

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 1
	}
	if index == 1 {
		return 1
	} else {
		return (Fibonacci(index-2) + Fibonacci(index-1))
	}
}

func main() {
	output := Fibonacci(6)
	fmt.Println(output)
}
