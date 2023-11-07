package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	program_name := os.Args[0]

	series := []rune(program_name)
	for _, name := range series[2:] { //  C: we made 2:0 so this part is not visible // C: bu kısmın gözükmemesi için 2:0 yaptık
		//z01.PrintRune(name)
		fmt.Print(string(name))
	}
	z01.PrintRune('\n')
}
