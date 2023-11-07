package main

import (
	"fmt"
	"os"
	"strconv"
)

//Bir int içine sığan iki tam pozitif tamsayıyı temsil eden iki dize alan bir program yazın.Program bunların en büyük ortak bölenini ve ardından bir satırsonu ('\n') görüntüler.
//Write a program that takes two string representing two strictly positive integers that fit in an int.The program displays their greatest common divisor followed by a newline ('\n').

func findGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}

	num1, err1 := strconv.Atoi(os.Args[1])
	num2, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil || num1 <= 0 || num2 <= 0 {
		fmt.Println()
		return
	}

	gcd := findGCD(num1, num2)

	fmt.Println(gcd)
}
