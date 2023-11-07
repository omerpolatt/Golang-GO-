package main

import "fmt"

// veri yapılarını (dilimleri, kanallar ve haritalar gibi) oluşturmak ve başlatmak için kullanılır
//used to create and initialize data structures (such as slices, channels and maps)

func MakeRange(min, max int) []int {
	if min >= max {
		var tab []int = nil
		return tab
	} else {
		tab := make([]int, max-min)
		j := 0
		for i := min; i < max; i++ {
			tab[j] = i
			j++
		}
		return tab
	}
}

func main() {
	output := MakeRange(1, 5)
	fmt.Println(output)
}
