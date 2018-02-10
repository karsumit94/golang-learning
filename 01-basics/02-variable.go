package main

import "fmt"
import "time"

func main() {
	name := "Sumit"
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := `Broadsoft`
	f := 'A'

    var g int
	var h string
	var i float64
	var j bool
	var k string
	var l time.Time

	var m,n int
	var o,p,q int = 1,2,3
	var r,s,t = 2,true,"PacketSmart"

	fmt.Println("Hello ", name)
	fmt.Println()
	fmt.Println()
	
	fmt.Printf("%v \t  %T    \n", a,a)
	fmt.Printf("%v \t  %T    \n", b,b)
	fmt.Printf("%v \t  %T    \n", c,c)
	fmt.Printf("%v \t  %T    \n", d,d)
	fmt.Printf("%v \t  %T    \n", e,e)
	fmt.Printf("%v \t  %T    \n", f,f)
	fmt.Println()
	fmt.Println()

	fmt.Printf("%v \t  %T    \n", g,g)
	fmt.Printf("%v \t  %T    \n", h,h)
	fmt.Printf("%v \t  %T    \n", i,i)
	fmt.Printf("%v \t  %T    \n", j,j)
	fmt.Printf("%v \t  %T    \n", k,k)
	fmt.Printf("%v \t  %T    \n", l,l)
	fmt.Println()
	fmt.Println()

	fmt.Printf("%v \t  %T    \n", m,m)
	fmt.Printf("%v \t  %T    \n", n,n)
	fmt.Printf("%v \t  %T    \n", o,o)
	fmt.Printf("%v \t  %T    \n", p,p)
	fmt.Printf("%v \t  %T    \n", q,q)
	fmt.Printf("%v \t  %T    \n", r,r)
	fmt.Printf("%v \t  %T    \n", s,s)
	fmt.Printf("%v \t  %T    \n", t,t)
	fmt.Println()
    fmt.Println(b, o, p, q)


}
