package main

import "fmt"

const (
	a = iota // 0
	b        // 1
	c = iota * 10 // 2  * 10 = 20
)

const (
	d = iota // 0
	e        // 1
	f        // 2
)

const (
	_  = iota             // 0\t\t\t\t
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	fmt.Println("Unit\tBinary\t\t\t\t\t\tDecimal")
	fmt.Printf("KB\t")
	fmt.Printf("%b\t\t\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("MB\t")
	fmt.Printf("%b\t\t\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("GB\t")
	fmt.Printf("%b\t\t\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("TB\t")
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
}

// Go's iota identifier is used in const declarations to 
// simplify definitions of incrementing numbers. Because 
// it can be used in expressions, it provides a generality 
// beyond that of simple enumerations.
// Iota: http://golang.org/doc/go_spec.html#Iota

// Constant declarations: http://golang.org/doc/go_spec.html#Constant_declarations