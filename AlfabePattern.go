package main

import "fmt"

// Write a program that displays the alphabet, with even letters in uppercase, and odd letters in lowercase, followed by a newline ('\n').
//Alfabeyi çift harfler büyük, tek harfler küçük olacak şekilde görüntüleyen ve ardından satırsonu ('\n') gelen bir program yazın.

func Alfabe() {
	sayac := 0
	for i := 'a'; i <= 'z'; i++ {
		if sayac%2 == 0 {
			fmt.Print(string(i))
		} else {
			fmt.Print(string(i - 32)) // Küçük harfi büyük harfe dönüştür
		}
		sayac++
	}
	fmt.Println()
}

func main() {
	Alfabe()
}
