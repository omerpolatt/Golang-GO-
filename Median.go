package main

import "fmt"

func Median(a, b, c, d, e int) int {
	numbers := []int{a, b, c, d, e}
	// Diziyi küçükten büyüğe sırala
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i] // burada sayıların yerini değiştirmiş oluyoruz büyük küçük yer değişyiriyorum. // buble sort
			}
		}
	}
	// Ortadaki değeri döndür // median value
	return numbers[2]
}

func main() {
	middle := Median(2, 3, 8, 5, 7)
	fmt.Println(middle)
}
