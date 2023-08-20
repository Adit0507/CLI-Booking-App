package main
import "fmt"

func main(){
	// := only applies to variables
	conferenceName :="Go Conference"
	const conferenceTickets int = 70
	var remainingTickets uint = 50

	// checking the data types of the variables
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are left", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

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

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets, You will receive confirmation mail at %v\n", firstName, lastName, userTickets,email)

	fmt.Printf("%v remaining Tickets for %v\n", remainingTickets, conferenceName)
}
