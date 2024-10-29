package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int32 = 50
	var remTickets int32 = 50

	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and currently available are %v", conferenceTickets, remTickets)
	fmt.Println("Get your tickets here!!")

	// var userName; //Go cannot infer the dataType so we need explicit data type declaration
	var userName string
	var emailID string
	var ticketBought int // use int as it auto detects the OS if it is 32 or 64

	// fmt.Scan(userName) but doesn't take an input
	// therefore we need to pass it by reference like:
	// fmt.Scan(&userName)
	fmt.Scanf("%s", &userName)
	fmt.Scan(&emailID)
	fmt.Scan(&ticketBought)
	remTickets -= ticketBought
	fmt.Println("Hello", userName)          // by default there is space after "Hello"
	fmt.Println("Your emailId is", emailID) // by default there is space after "Hello"
	fmt.Println("Successfully bought", ticketBought, "tickets")
}
