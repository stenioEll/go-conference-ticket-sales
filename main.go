package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName string = "Go Amazing Conference" //name we will reference a lot
	const conferenceTickets int = 50                    //creating the tickets
	var remainingTickets uint = 50                      //dont accept negative values
	var bookings []string                               //list of bookings

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T.\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v ourbooking aplication.\n", conferenceName)
	fmt.Printf("We have total of: %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets!")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		//ask used for their name
		println("Please enter your first name:")
		fmt.Scan(&firstName)

		println("Please enter your last name:")
		fmt.Scan(&lastName)

		println("Please enter your email address:")
		fmt.Scan(&email)

		println("Please enter number of tickets:")
		fmt.Scan(&userTickets)

		//Validation of data

		isValidName := len(firstName) >= 2 && len(lastName) >= 2 //validation of names
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		//logic
		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
			fmt.Printf("We still have %v tickets for %v\n", remainingTickets, conferenceName)

			firstNames := []string{} //define a slice only for the first names
			for _, booking := range bookings {
				var names = strings.Fields(booking)       //Fields splits the string with white space as separator
				firstNames = append(firstNames, names[0]) //saving in our slice
			}

			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				//end of program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short ")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign ")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid ")
			}
			fmt.Println("Your input data is invald, try again")
			continue
		}

	}

}
