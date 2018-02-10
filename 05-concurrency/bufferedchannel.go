package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	//Send value to buffered channel
	messages <- "Chennai"
	messages <- "Kolkata"
	messages <- "Delhi" // This will give exception if the size is set to less than 3
	//Recieve value from buffered channel
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	messages <- "Gujarat"
	messages <- "Mumbai"
	//Recieve value from buffered channel
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
