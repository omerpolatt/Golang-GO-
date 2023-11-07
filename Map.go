package main

import "fmt"

//Map anahtar-değer çiftlerini saklayabilen bir veri yapısıdır.
//Map is a data structure that can store key-value pairs.

func main() {

	student := make(map[string]int)

	student["Alice"] = 95
	student["Bob"] = 88
	student["Charlie"] = 72

	fmt.Println(student["Alice"])

}
