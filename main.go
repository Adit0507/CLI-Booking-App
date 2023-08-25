package main
import (
	"fmt"
	"strings"
)

func main(){
	// := only applies to variables
	conferenceName :="Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	bookings:= []string{} 

	fmt.Printf("Welcome to %v booking application\nWe have total of %v tickets and %v tickets are left.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)

	for {
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
		
		isValidName := len(firstName) >= 2 && len(lastName) >=2
		isValidEmail :=	strings.Contains(email, "@")
	 	isValidUserTicketNumber := userTickets > 0 && userTickets <=  remainingTickets

		// imp. condition
		if isValidEmail && isValidName && isValidUserTicketNumber{
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " "+ lastName)
		
			fmt.Printf("Thank you %v %v for booking %v tickets, You will receive confirmation mail at %v\n", firstName, lastName, userTickets,email)
			fmt.Printf("%v remaining Tickets for %v\n", remainingTickets, conferenceName)
			
			firstNames := []string{}
			for _, booking := range bookings {
				var names =	strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
	
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
