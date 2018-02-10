package main

import "fmt"

func main() {

	//nested for loop
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, " - ", j)
		}
	}
	fmt.Println()
	fmt.Println()

	// for while loop
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println()
	fmt.Println()


	// for loop with condition 
	for i := 1; i < 5; i++ {
		if i%2 == 1 {
			fmt.Println(i, "is Odd")
		} else {
			fmt.Println(i, "is Even")
		}
	}
	fmt.Println()
	fmt.Println()

}