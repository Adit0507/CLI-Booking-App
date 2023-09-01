package main

import (
	"booking-app/helper"
	"fmt"
	"time"
	"sync"
)

// := only applies to variables
var conferenceName ="Go Conference"
const conferenceTickets uint = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0) 
 
type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
	isOptedInForNewsLetter bool
}

var wg = sync.WaitGroup{}

func main(){
	
	greetUsers()	
		firstName , lastName , email , userTickets := getUserInput()
  		isValidEmail, isValidName, isValidUserTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)		
		
		// imp. condition
		if isValidEmail && isValidName && isValidUserTicketNumber{

			bookTicket(userTickets, firstName, lastName, email)
			
			wg.Add(1)
			go sendTicket(userTickets,firstName, lastName, email)	//starts a new goroutine
			
			// call function print first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of the bookings are %v \n", firstNames)
	
			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				// end program
				fmt.Println("Our conference is booked out. See you next year 😊")
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
		wg.Wait()
}

func greetUsers(){
	fmt.Printf("Welcome to %v booking application", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are left.\nGet your tickets here to attend\n", conferenceTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
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

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets, You will receive confirmation mail at %v\n", firstName, lastName, userTickets,email)
	fmt.Printf("%v remaining Tickets for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(4*time.Second)
	var ticket=  fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("--------------------------")
	fmt.Printf("Sending ticket 🚀:\n%v \nto email address %v\n", ticket, email)
	fmt.Println("--------------------------")

	wg.Done()
}