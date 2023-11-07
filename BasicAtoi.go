package main

import "fmt"

func BasicAtoi(s string) int { // Our program that gives the numbers in string
	// Stringler içindeki rakamları çıktıya veren programımız
	sayi := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			sayi = sayi*10 + int(c-'0') // int() li kısmı  ascii değerden  kurtarmak için bunu yapıyoruz
		} // sayi*10 kısmını ise basamak basamak artırması için
	}
	return sayi
}

func main() {
	output := BasicAtoi("Hel23lo world123")
	fmt.Println(output)
}
