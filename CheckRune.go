package main

import "fmt"

func IsUpper(s string) bool { // Capitalization only
	//Yalnızca büyük harf kontrolü
	dizi := []rune(s)
	for i := 0; i <= (len(dizi))-1; i++ {
		if dizi[i] < 'A' || dizi[i] > 'Z' {
			return false
		}
	}
	return true
}

func IsLower(s string) bool { // Lowercase check only
	// Yalnızca küçük harf kontrolü
	dizi := []rune(s)
	for i := 0; i <= (len(dizi))-1; i++ {
		if dizi[i] < 'a' || dizi[i] > 'z' {
			return false
		}
	}
	return true
}

func IsNumeric(s string) bool { // Digit check only
	// Yalnızca rakam kontrolü
	dizi := []rune(s)
	for i := 0; i <= (len(dizi) - 1); i++ {
		if dizi[i] < '0' || dizi[i] > '9' {
			return false
		}
	}
	return true
}

func IsAlpha(s string) bool { // Special Character control
	//Özel Karakter kontrolü
	liste := []rune(s)
	for i := 0; i <= (len(liste) - 1); i++ {
		if (liste[i] < 'a' || liste[i] > 'z') && (liste[i] < 'A' || liste[i] > 'Z') && (liste[i] < '0' || liste[i] > '9') {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))

	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))

	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))

	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))

}
