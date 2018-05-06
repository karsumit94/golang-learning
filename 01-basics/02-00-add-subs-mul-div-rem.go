package main

import "fmt"

func main() {
	fmt.Println("Hello, world !!")

	a := 5 // Initiate a variable [ datatype is not mandatory ]
	b := 35
	c := 3
	sum := a + b
	mul := a * c
	div := b / a
	frac := float64(b) / float64(c)
	rem := b % c

	fmt.Printf("%v + %v = %v \n", a, b, sum)
	fmt.Printf("%v x %v = %v \n", a, c, mul)
	fmt.Printf("%v / %v = %v \n", b, a, div)
	fmt.Printf("%v / %v = %f \n", b, c, frac)
	fmt.Printf("%d %% %v = %v \n", b, c, rem)
	fmt.Println()
	fmt.Println()

}
