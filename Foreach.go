package main

import "fmt"

//The function specified with foreach will move over the elements and perform the operation in the given function on each element.
//Foreach ile belirtilen fonksiyon elemanlar üzerinde gezinip her elemana verilen fonksiyondaki işlemi yaptıracaktır

func ForEach(f func(int), a []int) {
	for _, element := range a {
		f(element)
	}
}

func PrintNbr(x int) {
	fmt.Print(x + 1)
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
}
