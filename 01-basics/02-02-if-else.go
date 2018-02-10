package main

import "fmt"

func main() {


	if true {
		fmt.Println("This ran")
	}

	if false {
		fmt.Println("This did not run")
	}
	fmt.Println()
	fmt.Println()
	
	if !true {
		fmt.Println("This did not run")
	}

	if !false {
		fmt.Println("This ran")
	}
	fmt.Println()
	fmt.Println()

    b := true
	if food := "Chocolate"; b {
		fmt.Println(food)
	}
	//fmt.Println(food)
	fmt.Println()
	fmt.Println()

	x := 13 % 3
	fmt.Println(x)
	if x == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
	fmt.Println()
	fmt.Println()

}