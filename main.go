package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

// := only applies to variables
var conferenceName ="Go Conference"
const conferenceTickets uint = 50
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0) 

func main(){
	
	greetUsers()

	for {
		firstName , lastName , email , userTickets := getUserInput()
  		isValidEmail, isValidName, isValidUserTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)		
		
		// imp. condition
		if isValidEmail && isValidName && isValidUserTicketNumber{

			bookTicket(userTickets, firstName, lastName, email)
			
			// call function print first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of the bookings are %v \n", firstNames)
	
			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				// end program
				fmt.Println("Our conference is booked out. See you next year 😊")
				break
			}
		} else {
			
			if !isValidName {
				fmt.Println("First name or last name u entered is too short")
			}

			if !isValidEmail {
				fmt.Println("email address doesnt contain @ sign")
			}

			if !isValidUserTicketNumber {
				fmt.Println("No. of tickets u entered is invalid")
			}
		}
	}
}

func greetUsers(){
	fmt.Printf("Welcome to %v booking application", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are left.\nGet your tickets here to attend\n", conferenceTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	
	return firstNames
}

func getUserInput() (string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	
	// Ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) //pointer
	
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	
	fmt.Println("Enter your email address")
	fmt.Scan(&email)
	
	fmt.Println("Enter no. of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint,firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create map for a user
	var userData = make(map[string] string)
	userData["firstName"]= firstName
	userData["lastName"] = lastName
	userData["email"] = email

	userData["number"]= strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets, You will receive confirmation mail at %v\n", firstName, lastName, userTickets,email)
	fmt.Printf("%v remaining Tickets for %v\n", remainingTickets, conferenceName)
	
}