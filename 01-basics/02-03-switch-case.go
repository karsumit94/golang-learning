package main

import "fmt"

//  switch on types
//  -- normally we switch on value of variable
//  -- go allows you to switch on type of variable

type contact struct {
	greeting string
	name     string
}
// SwitchOnType works with interfaces
func SwitchOnType(x interface{}) {
	switch x.(type) { // this is an assert; asserting, "x is of this type"
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case contact:
			fmt.Println("contact")
		default:
			fmt.Println("unknown")
	}
}

func main() {

	//switch case
    switch "1" {
		case "1", "2":
			fmt.Println("1 or 2")
		case "3":
			fmt.Println("3")
			fallthrough
		case "4", "5":
			fmt.Println("4 or 5")
		default:
			fmt.Println("Default")
	}
	fmt.Println()
	fmt.Println()


	myFriendsName := "Aman"
	switch {
		case len(myFriendsName) == 4:
			fmt.Println("Hi my friend with name of length 4")
		case myFriendsName == "Trisha":
			fmt.Println("Hi Trisha")
		case myFriendsName == "Daphne":
			fmt.Println("Hi Daphne")
		case myFriendsName == "Barath", myFriendsName == "Krupesh":
			fmt.Println("Your name is either Barath or Krupesh")
		case myFriendsName == "Ranjitha":
			fmt.Println("Hi Ranjitha")
		default:
			fmt.Println("nothing matched; this is the default")
	}
	/*
	  expression not needed
	  -- if no expression provided, go checks for the first case that evals to true
	  -- makes the switch operate like if/if else/else
	  cases can be expressions
	*/
	fmt.Println()
	fmt.Println()


	SwitchOnType(7)
	SwitchOnType("Sayandeep")
	var t = contact{"Good to see you,", "Guneet"}
	SwitchOnType(t)
	SwitchOnType(t.greeting)
	SwitchOnType(t.name)
}


