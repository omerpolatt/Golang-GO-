package main

import (
	"fmt"
	"os"
)

func Rot14(s string) string {
	str := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			str += string('a' + ((r - 'a' + 14) % 26))
		} else if r >= 'A' && r <= 'Z' {
			str += string('A' + ((r - 'A' + 12) % 26))
		} else {
			str += string(r)
		}
	}
	return str
}
func main() {
	args := os.Args[1]
	output := Rot14(args)
	fmt.Println(output)
}
