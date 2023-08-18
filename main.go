package main
import "fmt"

func main(){
	// := only applies to variables
	conferenceName :="Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 20

	// checking the data types of the variables
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)


	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are left", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// Ask user for their name
	var userName string
	var userTickets int

	userName = "Simon"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets", userName, userTickets)
}
