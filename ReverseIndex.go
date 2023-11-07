package main

import "fmt"

func ReverseMenuIndex(menu []string) []string {
	output := make([]string, len(menu))

	for i, order := range menu {
		output[len(menu)-i-1] = order
	}
	return output
}

func main() {

	fmt.Println(ReverseMenuIndex([]string{"desserts", "mains", "drinks", "starters"}))
}
