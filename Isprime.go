package main

import "fmt"

func IsPrime(nb int) bool {
	if nb <= 0 || nb == 1 {
		return false
	} else {
		sayac := 0
		for i := 2; i <= nb/2; i++ {
			if nb%i == 0 {
				sayac++
			}
		}
		if sayac == 0 {
			return true
		} else {
			return false
		}
	}
}

func main() {
	output := IsPrime(15)
	fmt.Println(output)
}
