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

type Director struct {
	Person
	Teams  []string
}

type Manager struct {
	Person
	Roles  []string
}

type Member struct {
	Person
	Skills []string
}

type User interface {
	PrintName()
	PrintDetails()
}

type Team struct {
	Name, Description string
	Users             []User
}

//A method that accepts person
func (p Person) PrintName() {
	fmt.Printf("\n%s %s\n", p.FirstName, p.LastName)
}

//A method that accepts person
func (p Person) PrintDetails() {
	fmt.Printf("[Date of Birth: %s, Email: %s, Location: %s ]\n", p.Dob.String(), p.Email, p.Location)
}

//overrides PrintDetails
func (a Manager) PrintDetails() {
	//Call person PrintDetails
	a.Person.PrintDetails()
	fmt.Println("Manager Roles:  ")
	for _, v := range a.Roles {
		fmt.Println(v)
	}
}
func (a Director) PrintDetails() {
	//Call person PrintDetails
	a.Person.PrintDetails()
	fmt.Println("Teams :  ")
	for _, v := range a.Teams {
		fmt.Println(v)
	}
}
//overrides PrintDetails
func (m Member) PrintDetails() {
	//Call person PrintDetails
	m.Person.PrintDetails()
	fmt.Println("Skills:")
	for _, v := range m.Skills {
		fmt.Println(v)
	}
}

func (t Team) GetTeamDetails() {
	fmt.Printf("Team: %s  - %s\n", t.Name, t.Description)
	fmt.Println("Details   of the team members:")
	for _, v := range t.Users {
		v.PrintName()
		v.PrintDetails()
	}
}

func main() {
	sjayaraman := Director{
		Person{
			"Shiva",
			"Jayaraman",
			time.Date(1970, time.January, 17, 0, 0, 0, 0, time.UTC),
			"sjayaraman@email.com",
			"Chennai"},
		[]string{"PacketSmart", "MAPL"},
	}
	vnallusamy := Manager{
		Person{
			"Valarmathi",
			"Nallusamy",
			time.Date(1984, time.June, 17, 0, 0, 0, 0, time.UTC),
			"vnallusamy@email.com",
			"Chennai"},
		[]string{"Manage Team", "Manage Tasks"},
	}
	skar := Member{
		Person{
			"Sumit",
			"Kar",
			time.Date(1994, time.March, 14, 0, 0, 0, 0, time.UTC),
			"skar@email.com",
			"Kolkata"},
		[]string{"Java", "C", "Go"},
	}
	ajain := Member{
		Person{
			"Aman",
			"Jain",
			time.Date(1994, time.August, 23, 0, 0, 0, 0, time.UTC),
			"ajain@email.com",
			"Murshidabad"},
		[]string{"Java", "DBMS", "C"},
	}
	krupeshpatel := Member{
		Person{
			"Krupesh",
			"Patel",
			time.Date(1993, time.May, 24, 0, 0, 0, 0, time.UTC),
			"krupeshpatel@email.com",
			"Gujarat"},
		[]string{"Shell", "Perl", "Python"},
	}
	team := Team{
		"PacketSmart",
		"PacketSmart Development Team",
		[]User{sjayaraman, vnallusamy, skar, ajain, krupeshpatel},
	}
	//get details of Team
	team.GetTeamDetails()
}
