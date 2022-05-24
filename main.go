package main

import (
	"fmt"
	"sync"
	"time"
)

// Package variables
var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	// Welcome screen of the cli
	greetUsers()

	// Asking user inputs
	firstName, lastName, emailAddr, userTickets := getUserInput()

	// Checking the validity of the user inputs
	isValidName, isValidTicketNumber, isValidEmail := ValidateUserInput(firstName, lastName, emailAddr, userTickets, remainingTickets)

	// Check the validity of the ticket user entered
	if isValidTicketNumber && isValidEmail && isValidName {
		// Booking tickets
		bookTicket(firstName, lastName, emailAddr, userTickets)
		// Sending tickets
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, emailAddr) // "Go" is a keyword for creating new threads for avoid the delays
		// Print only first names
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are %v \n", firstNames)

		var noTicketsRemaining bool = remainingTickets == 0

		// Ending program
		if noTicketsRemaining {
			fmt.Println("No tickets available, please come again next year")
		}
	} else {
		if !isValidName {
			fmt.Println("First name or last name you entered is too short")
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
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets")
	fmt.Println(remainingTickets, "tickets are still available")
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var emailAddr string
	var userTickets uint

	// Asking user inputs
	fmt.Println("Enter your first name ")
	fmt.Print(">> ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name ")
	fmt.Print(">> ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address ")
	fmt.Print(">> ")
	fmt.Scan(&emailAddr)

	fmt.Println("Enter number of tickets ")
	fmt.Print(">> ")
	fmt.Scan(&userTickets)

	return firstName, lastName, emailAddr, userTickets
}

func bookTicket(firstName string, lastName string, emailAddr string, userTickets uint) {
	remainingTickets -= userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           emailAddr,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of booking is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. \n", firstName, lastName, userTickets)
	fmt.Printf("You will receive confirmation e-mail at %v\n", emailAddr)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second) // Setting a delay for testing "go" keyword for multi threading
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("#############")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#############")
	wg.Done()
}
