package main

import (
	"fmt"
	"strings"
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	seriesname := "Go Conference"
	const Tickets = 50
	var remainingTickets uint = 50
	bookings := make([]UserData, 0) // Correctly initialize the bookings slice

	fmt.Printf("Welcome to the %v booking application\n", seriesname)
	fmt.Printf("We have a total of %v tickets and %v are still available to purchase\n", Tickets, remainingTickets)
	fmt.Println("Get your tickets here now to attend today!")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Please enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email address:")
		fmt.Scan(&email)

		fmt.Println("Please enter the number of tickets you want to buy:")
		fmt.Scan(&userTickets)

		// User input validation
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".com")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets -= userTickets

			// Creating a userData struct
			var userData = UserData{
				firstName:       firstName,
				lastName:        lastName,
				email:           email,
				numberOfTickets: userTickets,
			}

			// Store userData struct in bookings
			bookings = append(bookings, userData)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v shortly.\n", firstName, lastName, userTickets, email)
			fmt.Printf("There are now %v tickets remaining for the %v.\n", remainingTickets, seriesname)

			firstNames := []string{}
			for _, booking := range bookings {
				firstNames = append(firstNames, booking.firstName)
			}
			fmt.Printf("The first names of all our bookings: %v\n", firstNames)

			sendTicket(userTickets, firstName, lastName, email)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Invalid name entered. First and last name must be two characters or more. Try again.")
			}
			if !isValidEmail {
				fmt.Println("Invalid email input. Try again.")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is invalid. Try again.")
			}
		}
	}
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#############")
	fmt.Printf("Sending %v\n \nto email address %v\n", ticket, email)
	fmt.Println("############")
}
