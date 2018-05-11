package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	x := rand.Intn(300000)
	fmt.Println(x)
}
