package main

import "fmt"

func ToLower(s string) string { // Convert Uppercase to Lowercase 
	liste := []rune(s)
	for i := range liste {
		if liste[i] >= 'A' && liste[i] <= 'Z' {
			liste[i] = liste[i] + 32
		}
	}
	return string(liste)
}

func ToUpper(s string) string { // Convert Lowercase to Uppercase 
	liste := []rune(s)
	for i := 0; i <= (len(liste) - 1); i++ {
		if liste[i] >= 'a' && liste[i] <= 'z' {
			liste[i] = liste[i] - 32
		}
	}
	return string(liste)
}

func main() {
	output1 := ToLower("Hello! How are you?")
	output2 := ToUpper("Hello! How are you?")

	fmt.Println(output1)
	fmt.Println(output2)
}
