package main

import (
	"fmt"
	"os"
)

func canWriteWithOrder(str1, str2 string) bool {
	// İlk dize (str1) ve ikinci dize (str2) için iki indeks tanımla.
	index1, index2 := 0, 0

	// İlk dizeyi (str1) ikinci dizede (str2) sırasıyla yazabilir miyiz kontrol etmek için bir döngü başlatılır.
	// İlk dizeyi gezme indeksi (index1) ve ikinci dizeyi gezme indeksi (index2) kullanılır.

	// Her iki dize de sona ulaşmadan ve karakterler eşleşmeye devam ettikçe döngü devam eder.
	for index1 < len(str1) && index2 < len(str2) {
		// Eğer şu anki karakterler eşleşiyorsa, ilerlememiz gereken ilk dizide bir sonraki karakteri gösterir (index1 arttırılır).
		if str1[index1] == str2[index2] {
			index1++
		}
		// İkinci dizide her durumda bir sonraki karaktere geçeriz (index2 arttırılır).
		index2++
	}

	// Eğer ilk dizeyi tamamen ikinci dizede yazabildiysek, bu fonksiyon `true` döndürür.
	return index1 == len(str1)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	if canWriteWithOrder(str1, str2) {
		fmt.Println(str1)
	} else {
		fmt.Println()
	}
}
