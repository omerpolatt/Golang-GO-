package main

import "fmt"

func StrLen(s string) int { // Function that counts all characters in a String
	sayac := 0
	for range s {
		sayac++
	}
	return sayac
}

func AlphaCount(s string) int { // Our Function that counts the letters in String
	count := 0
	for _, yazi := range s {
		if (yazi <= 'z' && yazi >= 'a') || (yazi <= 'Z' && yazi >= 'A') {
			count++
		}
	}
	return count
}

func main() {
	output1 := StrLen("Hello 78 World!    4455 /")
	output2 := AlphaCount("Hello 78 World!    4455 /")

	fmt.Println(output1)
	fmt.Println(output2)
}
