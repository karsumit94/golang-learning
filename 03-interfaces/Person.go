package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName, LastName string
	Dob                 time.Time
	Email, Location     string
}

//A person method with pointer receiver
func (p *Person) PrintName() {
	fmt.Printf("\n%s %s\n", p.FirstName, p.LastName)
}

//Another person method with pointer receiver
func (p *Person) PrintDetails() {
	fmt.Printf("[Date of Birth: %s, Email: %s, Location: %s ]\n", p.Dob.String(), p.Email, p.Location)
}

//A person method with pointer receiver
func (p *Person) ChangeLocation(newLocation string) {
	p.Location = newLocation
}
func main() {
	p := &Person{
		"Sumit",
		"Kar",
		time.Date(1994, time.March, 14, 0, 0, 0, 0, time.UTC),
		"skar@email.com",
		"Kolkata",
	}
	p.PrintName()
	p.PrintDetails()
	time.Sleep(2 * time.Second)
	fmt.Printf("\nRelocating to new location.\n\n")
	time.Sleep(2 * time.Second)
	p.ChangeLocation("Chennai")
	p.PrintName()
	p.PrintDetails()

}
