package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//by default the seed for rand in golang is 1 
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(300000)
	fmt.Println(x)
}
