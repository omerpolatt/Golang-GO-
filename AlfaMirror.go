package main

import (
	"fmt"
	"os"
)

//Write a program called alphamirror that takes a string as argument and displays this string after replacing each alphabetical character with the opposite alphabetical character. The case of the letter stays the same
//Bir dizeyi argüman olarak alan ve her alfabetik karakteri karşıt alfabetik karakterle değiştirdikten sonra bu dizeyi görüntüleyen alphamirror adlı bir program yazın. Harfin büyük/küçük harf durumu aynı kalır,

func alphaMirror(s string) string {
	result := ""
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') { //letter control condition // harf kontrol koşulu
			if char >= 'a' && char <= 'z' {
				result += string('z' - (char - 'a'))
			}
			if char >= 'A' && char <= 'Z' {
				result += string('Z' - (char - 'A'))
			}
		} else { // harf dışındaki karakterler için olan kısım  //the part for characters other than letters
			result += string(char)
		}
	}

	return result
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
	} else {
		input := os.Args[1]
		result := alphaMirror(input)
		fmt.Println(result)
	}
}
