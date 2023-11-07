package main

import (
	"fmt"
	"os"
)

//İlk argümanı bir harfin (2. argüman) başka bir harfle (3. argüman) değiştirileceği bir dize olan 3 argüman alan bir program yazın.
//Eğer argüman sayısı 3'ten farklıysa, program bir satırsonu ('\n') görüntüler.
//İkinci bağımsız değişken ilk bağımsız değişken (dize) içinde yer almıyorsa, program dizeyi yeniden yazar ve ardından bir satırbaşı ('\n') görüntüler.

//Write a program that takes 3 arguments, the first argument is a string in which to replace a letter (the 2nd argument) by another one (the 3rd argument).
//If the number of arguments is different from 3, the program displays a newline ('\n').
//If the second argument is not contained in the first one (the string) then the program rewrites the string followed by a newline ('\n').

func replaceChars(input, findChar, replaceChar string) string {
	kelime := make([]byte, len(input)) // replaced adlı bir dizi oluşturduk girilen girdi uzunluğunda .
	for i := 0; i < len(input); i++ {
		if string(input[i]) == findChar { // i deki karakter istenilen karakterle aynı ise
			kelime[i] = replaceChar[0]
		} else { // findchar ile erişmiyorsa
			kelime[i] = input[i]
		}
	}
	return string(kelime)
}

func main() {
	if len(os.Args) != 4 { // fazla argüman girdiğin zaman boş dönmesini istediğimiz durum
		fmt.Println()
		return
	}

	input := os.Args[1] // burada os ile aldığımız argümanları belirliyoruz  fonksiyonumuzda argümann  olarka verilenlerle eşleştiriryoruz
	findChar := os.Args[2]
	replaceChar := os.Args[3]

	result := replaceChars(input, findChar, replaceChar) // fonksiyonumuzu main içinde çalıştırıp ekrana bastırlmasını tamamladık .
	fmt.Println(result)
}
