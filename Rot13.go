package main

import (
	"fmt"
	"os"
)

//Bir dizeyi alan ve her bir harfini alfabetik sıraya göre 13 boşluk ilerideki harfle değiştirerek görüntüleyen bir program yazın. Harf büyüklüğü etkilenmez.

//Write a program that takes a string and displays it, replacing each of its letters by the letter 13 spaces ahead in alphabetical order. Case remains unaffected.

func Rot13(s string) string {
	str := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			str += string('a' + ((r - 'a' + 13) % 26))
		} else if r >= 'A' && r <= 'Z' {
			str += string('A' + ((r - 'A' + 13) % 26))
		} else {
			str += string(r)
		}
	}
	return str
}
func main() {
	args := os.Args[1]
	output := Rot13(args)
	fmt.Println(output)
}
