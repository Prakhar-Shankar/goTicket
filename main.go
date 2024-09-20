package main

import "fmt"

func main(){

	var conferenceName = "Go Conference"
	const conferenceTicket = 50
	var remainingTicket uint = 50
	var bookings [50] string // Defining an array of 50 strings. Yaha par bhi curly braces use krke value assing kr skte hain lekin we are assinging the values later on.

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets out of which %v tickets are available\n", conferenceTicket, remainingTicket)
	fmt.Println("You can book your tickets here to attend", conferenceName)

	var firstName string
	var lastName string 
	var email string 
	var userTicket uint 

	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter you email Id: ")
	fmt.Scan(&email)

	fmt.Println("Please enter the number of tickets you want to book: ")
	fmt.Scan(&userTicket)

	remainingTicket = remainingTicket - userTicket

	fmt.Printf("Thank you %v %v for booking %v tickets of the conference. You will shortly receive a confirmation email on %v\n 	", firstName, lastName, userTicket, email)
	bookings[0]= firstName+ " "+ lastName
	fmt.Printf("The first Value: %v\n", bookings[0])
	fmt.Printf("Now we have %v tickets left for the conference.\n", remainingTicket)
}