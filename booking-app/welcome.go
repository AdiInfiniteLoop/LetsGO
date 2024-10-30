package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int32 = 50
	remTickets := 50
	// Sizes in go have fixed length
	bookings := [50]string{"Aditya Pradhan", "Adi"}
	// var bookingsAlternative [50]string
	// var emptyArray []string{}
	// use append to make a dynamic-sized array. E.g., arr.append(43)

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
	remTickets = remTickets - ticketBought
	fmt.Println("Hello", userName)          // by default there is space after "Hello"
	fmt.Println("Your emailId is", emailID) // by default there is space after "Hello"
	fmt.Println("Successfully bought", ticketBought, "tickets")
	bookings[0] = userName

	/*
		Array printing
	*/
	fmt.Printf("The whole bookings array is %v", bookings)
	fmt.Printf("The first element of the array is %v\n", bookings[0])
	fmt.Printf("The length of the bookings array is %v\n", len(bookings))
	fmt.Printf("The Type of the bookings array is %T\n", bookings)
	// array types in go are defined by [size]<type>
}
