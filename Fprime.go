package main

import (
	"fmt"
)

//Pozitif bir int alan ve asal çarpanlarını görüntüleyen ve ardından bir satırbaşı ('\n') getiren bir program yazın.
//Write a program that takes a positive int and displays its prime factors, followed by a newline ('\n').

func primeFactors(n int) {
	fmt.Print("Bölenler: ")
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			if isPrime(i) {
				fmt.Print(i)
				if n != i {
					fmt.Print(" ")
				}
			}
		}
	}
	fmt.Println()
}

func isPrime(nb int) bool {
	if nb <= 0 || nb == 1 {
		return false
	} else {
		count := 0

		for i := 2; i <= nb/2; i++ {
			if nb%i == 0 {
				count++
			}
		}
		if count == 0 {
			return true
		} else {
			return false
		}

	}
}

func main() {
	primeFactors(78)
}
