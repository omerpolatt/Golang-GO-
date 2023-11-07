package main

import "fmt"

func Unmatch(a []int) int {
	for _, i := range a {
		sayac := 0
		for _, x := range a {
			if i == x {
				sayac++
			}
		}
		if sayac == 1 || sayac%2 == 1 {
			return i
		}
	}
	return -1
}

func main() {

	slice := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := Unmatch(slice)
	fmt.Println(unmatch)
}
