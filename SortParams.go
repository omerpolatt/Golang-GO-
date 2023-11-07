package main

import (
	"os"

	"github.com/01-edu/z01"
)

func BubbleSort(array []string) []string {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func main() {
	for _, arg := range BubbleSort(os.Args[1:]) {
		for _, char := range arg {
			z01.PrintRune(rune(char))
		}
		z01.PrintRune('\n')
	}
}

/*
OUTPUT :
 go run .\SortParams.go 1 a 2 A 3 b 4 C
1
2
3
4
A
C
a
b
*/
