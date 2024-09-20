package main

import (
	"fmt"
	"strings"
)

func main(){

	var conferenceName = "Go Conference"
	const conferenceTicket = 50
	var remainingTicket uint = 50
	bookings:= []string{} // Defining an array of 50 strings. Yaha par bhi curly braces use krke value assing kr skte hain lekin we are assinging the values later on.
	// here instead of the array hamne slice use kar liya hai, slice ka use ham array ko dynamic banane ke liye karte hain. Slice ka use karne se ham array ki length ko badha sakte hain ya ghatasakte hain.

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets out of which %v tickets are available\n", conferenceTicket, remainingTicket)
	fmt.Println("You can book your tickets here to attend", conferenceName)

	
	for{

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
		bookings = append(bookings, firstName + " " + lastName) //append function use krke ham slice mein elements add kar skte hain.

		fmt.Printf("Thank you %v %v for booking %v tickets of the conference. You will shortly receive a confirmation email on %v\n 	", firstName, lastName, userTicket, email)
		
		fmt.Printf("Now we have %v tickets left for the conference.\n", remainingTicket)

		firstNames := []string{}
		for _, booking := range bookings{
			var names = strings.Fields(booking) // Fields ka use separate krne ke liye hota hai, for example agar hamare pas first name and last name tha toh ham chahte the to use only the first name uske liye hame use separate krna pdta, Fields function "strings" package se aaya hai.
			firstNames= append(firstNames, names[0])
		}


		fmt.Printf("The first names of bookings are: %v\n", firstNames)
	}

	
}