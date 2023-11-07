package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	parameters := os.Args
	for i := len(parameters) - 1; i > 0; i-- {
		for _, para := range parameters[i] {
			fmt.Print(string(para))
			//z01.PrintRune(para)
		}
		z01.PrintRune('\n')
	}
}

/*
 OUTPUT :
 go run .\PrintRevParams.go 5 "wed" 6 "e"
e
6
wed
5
*/
