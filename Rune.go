package main

import "fmt"

func FirstRune(s string) rune { // Printing the First Character
	liste := []rune(s)
	return liste[0]
}

func LastRune(s string) rune { // Printing the Last Character
	liste := []rune(s)
	for i := 0; i <= len(liste); i++ {
		return rune(liste[len(liste)-1])
	}
	return 0
}

func NRune(s string, n int) rune { // Print the entered number of characters
	liste := []rune(s)
	if n > len(liste) {
		return 0
	} else {
		for i := 1; i <= len(liste); i++ {
			if i == n {
				return rune(liste[i-1])
			}
		}
	}
	return 0
}

func main() {
	output1 := FirstRune("Hello!")
	output2 := LastRune("Hello!")
	output3 := NRune("Hello!", 3)

	fmt.Println(string(output1))
	fmt.Println(string(output2))
	fmt.Println(string(output3))
}
