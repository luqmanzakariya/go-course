package main

import "fmt"

func main () {
	// Example of variable declaration
	conferenceName := "Go conference"
	const conferenceTicket = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceName is %T conferenceTicket is %T remainingTickets is %T \n", conferenceName, conferenceTicket, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTicket, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")
	
	/*
	  # Using Println
		fmt.Println("Welcome to", conferenceName, "booking application")
		fmt.Println("We have total of", conferenceTicket, "tickets and", remainingTickets, "are still available.")
		fmt.Println("Get your tickets here to attend")
	*/

	var userName string
	var userTickets int
	// ask the user for their name
	
	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}