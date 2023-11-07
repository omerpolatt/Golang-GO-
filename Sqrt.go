package main

import "fmt"

// Girilen Sayının karakökü tam sayı ise o değeri döndüren , Karakökü tam sayıya denk gelmeyen sayılar içinde 0 değerini döndüren Fonksiyonumuz

// Our Function that returns that value if the square root of the entered number is an integer, and returns 0 for numbers whose square root does not correspond to an integer.
func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	sqrt := 1
	for sqrt*sqrt < nb {
		sqrt++
	}
	if sqrt*sqrt == nb {
		return sqrt
	}
	return 0
}

func main() {

	output := Sqrt(64)
	fmt.Print(output)

}
