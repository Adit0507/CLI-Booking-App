package main
import "fmt"

func main(){
	var conferenceName ="Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 20

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are left", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
