package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args[1:]

	for _, s := range a {
		for _, char := range s {
			if char >= 'a' && char <= 'z' {
				z01.PrintRune(char - 32)
			} else if char >= 'A' && char <= 'Z' {
				z01.PrintRune(char + 32)
			} else { // özel karakterler için olan durumda burası çalışacak
				z01.PrintRune(char)
			}
		}
		z01.PrintRune('\n')
	}
}
