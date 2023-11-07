package main

import "fmt"

func IterativeFactorial(nb int) int {
	factorial := 1
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	if nb < 25 {
		for i := 1; i <= nb; i++ {
			factorial = factorial * i
		}
		return factorial
	} else {
		return 0
	}
}

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	if nb < 25 {
		return (RecursiveFactorial(nb - 1)) * nb
	} else {
		return 0
	}
}

func main() {
	outpu1 := IterativeFactorial(5)
	outpu2 := RecursiveFactorial(2)

	fmt.Println(outpu1)
	fmt.Println(outpu2)
}
