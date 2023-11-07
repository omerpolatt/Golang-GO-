package main

import "fmt"

//Struct  veri tipi, farklı türlerdeki verileri bir arada saklamak için kullanılan bir veri yapısıdır.
//The Struct data type is a data structure used to store different types of data together.

type food struct {
	preptime int // struct özelliği
}

func FoodDeliveryTime(order string) int {
	var burger food
	burger.preptime = 15
	var chips food
	chips.preptime = 10
	var nuggets food
	nuggets.preptime = 12
	if order == "burger" {
		return burger.preptime
	}
	if order == "chips" {
		return chips.preptime
	}
	if order == "nuggets" {
		return nuggets.preptime
	}
	return 404 // üçünden birine eşit değilse hata yolla
}

func main() {
	fmt.Println(FoodDeliveryTime("burger"))
	fmt.Println(FoodDeliveryTime("chips"))
	fmt.Println(FoodDeliveryTime("nuggets"))
	fmt.Println(FoodDeliveryTime("burger") + FoodDeliveryTime("chips") + FoodDeliveryTime("nuggets"))

}
