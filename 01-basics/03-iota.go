package main

import "fmt"

const (
	a = iota // 0
	b        // 1
	c        // 2
)

const (
	d = iota // 0
	e        // 1
	f        // 2
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

// Go's iota identifier is used in const declarations to 
// simplify definitions of incrementing numbers. Because 
// it can be used in expressions, it provides a generality 
// beyond that of simple enumerations.
// Iota: http://golang.org/doc/go_spec.html#Iota

// Constant declarations: http://golang.org/doc/go_spec.html#Constant_declarations