package main

import "fmt"

//Bir dizi , dilim gibi veri tiplerine eleman eklemeye yarayan fonksiyonumuz
// Our function for adding elements to data types such as an array, slice

func AppendRange(min, max int) []int {
	if min >= max {
		var tab []int = nil
		return tab
	} else {
		tab := []int{}
		for i := min; i < max; i++ {
			tab = append(tab, i)
		}
		return tab
	}
}

func main() {
	output := AppendRange(5, 14)
	fmt.Print(output)
}
