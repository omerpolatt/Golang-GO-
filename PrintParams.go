package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

// Our function that prints the parameters entered when using a function

func main() {
	for _, parameters := range os.Args[1:] {
		for _, i := range parameters {
			//z01.PrintRune(i)
			fmt.Print(string(i))
		}
		z01.PrintRune('\n')
	}
}

/* OUTPUT :
go run .\PrintParams.go 5 "wed" 6 "e"
5
wed
6
e
*/
