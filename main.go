package main

import (
	"fmt"
	"sync"
	// "strings"
	"go-course/helper"
	// "strconv"
	"time"
)

// Example of variable declaration
const conferenceTicket int = 50
var conferenceName = "Go conference"
var remainingTickets uint = 50
// list of slice
// var bookings = []string{}
// list of map
// var bookings = make([]map[string]string, 0)
// list of struct
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main () {
	// fmt.Printf("conferenceName is %T conferenceTicket is %T remainingTickets is %T \n", conferenceName, conferenceTicket, remainingTickets)
	
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

		#Check array
		fmt.Printf("The whole array %v\n", bookings)
		fmt.Printf("The first value %v\n", bookings[0])
		fmt.Printf("Array type: %T\n", bookings)
		fmt.Printf("Array length: %v\n", len(bookings))
	*/

	/*
		# Declaring An Slice
		var bookings []string // example 1
		var bookings = []string{} // example 2
		bookings := []string{} // example 3

		# Check slices
		fmt.Printf("The whole Slice %v\n", bookings)
		fmt.Printf("The first value %v\n", bookings[0])
		fmt.Printf("Slice type: %T\n", bookings)
		fmt.Printf("Slice length: %v\n", len(bookings))	
	*/

	/*
		# Declaring a Boolean Variable
		var noTicketsRemaining bool = remainingTickets == 0
		noTicketsRemaining := remainingTickets == 0
	*/

	/*
		# Switch Statement

		city := "Hongkong"

		switch city {
			case "New York":
				fmt.Println("Execute code for booking New York conference tickets")
			case "Singapore":
				fmt.Println("Execute code for booking Singapore conference tickets")
			case "London", "Berlin":
				fmt.Println("Execute code for London and Berlin conference tickets")
			case "Mexico City", "Hongkong":
				fmt.Println("Execute code for booking Mexico City and Hongkong conference tickets")
			default:
				// Default handles the case, if no match is found
				fmt.Println("No valid city selected")
		}
	*/
	greetUsers()

	/*
	// === Infinite loop ===
	for {
		// Scan and get user input
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)
	
			// Call function print first names
			firstNames := getFirstName()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
	
			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year. ")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}
	*/

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		// Call function print first names
		firstNames := getFirstName()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year. ")
		}
	} else {
		if !isValidName {
			fmt.Println("First name or last name is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid")
		}
	}

	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTicket, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")
}

func getFirstName() []string {
	var firstNames = []string{}
	/*
	booking as map
	for _,booking := range bookings {
			// Using slice
			name := strings.Fields(booking)[0]
		firstNames = append(firstNames, booking["firstName"])
	}
	*/
	
	for _,booking := range bookings {
		firstNames = append(firstNames,booking.firstName)
	}
	
	return firstNames
}

func getUserInput()(string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
} 

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// Create a map for a user
	// var userData = make(map[string]string)
	/*
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	*/

	// Create a struct for user
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	// using slice
	// bookings = append(bookings, firstName + " " + lastName)

	// using list of map
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket (userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second) // delay 10 seconds
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###################")
	fmt.Printf("Sending ticket to \n %v \nto email address %v\n", ticket, email)
	fmt.Println("###################")
	wg.Done()
}