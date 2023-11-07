package main

import "fmt"

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	} else {
		carpim := 1
		for i := 1; i <= power; i++ {
			carpim = nb * carpim
		}
		return carpim
	}
}

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	} else {
		return nb * RecursivePower(nb, power-1)
	}
}

func main() {

	outpu1 := IterativePower(2, 5)
	outpu2 := RecursivePower(5, 2)
	fmt.Println(outpu1)
	fmt.Println(outpu2)
}
