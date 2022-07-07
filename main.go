package main

import "fmt"

func main () {
	// Example of variable declaration
	conferenceName := "Go conference"
	const conferenceTicket = 50
	var remainingTickets uint = 50
	var bookings []string

	// fmt.Printf("conferenceName is %T conferenceTicket is %T remainingTickets is %T \n", conferenceName, conferenceTicket, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTicket, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")
	
	/*
	  # Using Println
		fmt.Println("Welcome to", conferenceName, "booking application")
		fmt.Println("We have total of", conferenceTicket, "tickets and", remainingTickets, "are still available.")
		fmt.Println("Get your tickets here to attend")
	*/

	/*
		# ask the user for their name
		fmt.Scan(&firstName)
	*/

	/*
		# Pointer
		fmt.Println(remainingTickets) // print the value
		fmt.Println(&remainingTickets) // print the memory address of the variable
	*/

	/*
		# Declaring Array type
		var bookings = [50]string{} // Same with below
		var bookings [50]string // Same with above

		# Declaring Array type
		var bookings [50]string
		bookings[0] = "Nana"
	*/

	/*
		# Declaring An Slice
		var bookings []string // example 1
		var bookings = []string{} // example 2
		bookings := []string{} // example 3
	*/

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)
	
	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)

	/*
		// Check array
		fmt.Printf("The whole array %v\n", bookings)
		fmt.Printf("The first value %v\n", bookings[0])
		fmt.Printf("Array type: %T\n", bookings)
		fmt.Printf("Array length: %v\n", len(bookings))
	*/

	/*
		// Check slices
		fmt.Printf("The whole Slice %v\n", bookings)
		fmt.Printf("The first value %v\n", bookings[0])
		fmt.Printf("Slice type: %T\n", bookings)
		fmt.Printf("Slice length: %v\n", len(bookings))	
	*/

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("These are all our bookings: %v\n", bookings)
}