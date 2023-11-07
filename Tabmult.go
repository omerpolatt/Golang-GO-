package main

import (
	"fmt"
	"os"
	"strconv"
)

func tabMult(n int) {
	for i := 1; i <= 9; i++ {
		fmt.Printf("%d x %d = %d\n ", i, n, i*n)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Gecerli Argüman Giriniz")
		return
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	tabMult(number)
}
