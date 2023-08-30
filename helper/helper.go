package helper

import "strings"

// To export a func, just capitalize the first letter
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >=2
	isValidEmail :=	strings.Contains(email, "@")
	isValidUserTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidEmail, isValidUserTicketNumber, isValidName
}