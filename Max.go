package main

import "fmt"

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	maximum := a[0]
	for _, i := range a {
		if i > maximum {
			maximum = i
		}
	}
	return maximum
}

func main() {
	slice := []int{23, 123, 1, 11, 55, 93}
	max := Max(slice)
	fmt.Println(max)
}
