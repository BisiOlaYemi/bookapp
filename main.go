package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	conferenceName := "my Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUser()

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	for {

		var firstName string
		var lastName string
		var email string
		var userTicktets uint

		//ask user for their name
		fmt.Println("Enter your First Name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your Last Name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your Email Address:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTicktets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTicktets > 0 && userTicktets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTicktets
			bookings = append(bookings, firstName+""+lastName)
			go sendTicket(userTicktets, firstName, lastName, email)
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTicktets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			//call funtion to print first name
			printFirstname(bookings)

			if remainingTickets == 0 {
				//end Sales
				fmt.Println("Our Conference Ticket is sold out, Come back next year")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First Name or Last Name you entered is too short")
			}

			if isValidEmail {
				fmt.Println("The email address is incorrect")

			}

			if isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
}

func greetUser() {
	fmt.Println("Hi Good Morning!")

}

func printFirstname(bookings []string) {

	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("The first names of bookings are: %v\n", firstNames)

}

func sendTicket(userTicktets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicktets, firstName, lastName)
	fmt.Println("######")
	fmt.Printf("Sending ticket: \n %v \nto email address %v\n", ticket, email)
	fmt.Println("##########")
}
