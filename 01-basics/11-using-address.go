package main

import "fmt"

const metersToKm float64 = 0.001

func main() {
	var meters float64
	fmt.Print("Enter meters: ")
	fmt.Scan(&meters)
	kms := meters * metersToKm
	fmt.Println(meters, " meters is ", kms, " kms.")
}
