package main

import "fmt"

var x = 42
func increment() int {
	x++
	return x
}
func main() {
	//var x = 42
	increment()
	increment()
	fmt.Println(x)
	foo()

	{
		y := "Hello, World !!"
		fmt.Println(y)
	}
	// fmt.Println(y) // outside scope of y
	/*
	closure helps us limit the scope of variables used by multiple functions
	without closure, for two or more funcs to have access to the same variable,
	that variable would need to be package scope
	*/
}

func foo() {
	fmt.Println(x)
}


