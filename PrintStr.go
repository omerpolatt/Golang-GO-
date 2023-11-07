package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {

	for _, i := range s {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')

}

func StrRev(s string) string {

	empty := ""
	word := []rune(s)
	for i := len(s) - 1; i >= 0; i-- {
		empty += string(word[i])
	}
	return empty

}

func main() {
	PrintStr("Hello World!")

	s := StrRev("Hello World!")
	fmt.Println(s)
}
