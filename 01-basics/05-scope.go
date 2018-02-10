package main

import "fmt"

var x = 42

func main() {
	//var x = 42
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}
