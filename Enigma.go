package main

func Enigma(a ***int, b *int, c *******int, d ****int) {

	//Program to change the values of variables whose addresses are held with pointers between each other

	/*
		A system that takes pointers as arguments and uses them as
		a into c.
		c into d.
		d into b.
		b into a. program that makes value changes in the form
	*/

	x := *******c
	*******c = ***a
	y := ****d
	****d = x
	z := *b
	*b = y
	***a = z

}
