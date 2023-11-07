package main

import "fmt"

func JumpOver(str string) string {
	var out string
	for i := 2; i < len(str); i += 3 {
		out += string(str[i])
	}
	out += "\n"
	return out
}

func main() {

	fmt.Print(JumpOver("1010101010"))
	fmt.Print(JumpOver("t w e l v e"))
}
