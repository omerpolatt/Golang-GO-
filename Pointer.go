package main

import "fmt"

func main() {

	a := 23
	b := "Uniwoorkhub"

	a_adres := &a // adres tutucu pointer
	b_adres := &b

	fmt.Println(a_adres)
	fmt.Println(b_adres)

	a_degeri := *a_adres //deger gösterici
	b_degeri := *b_adres

	fmt.Println(a_degeri)
	fmt.Println(b_degeri)
}
