package main

import "fmt"

func main() {
	fmt.Println("Welcome users to our Booking App")
	string1 := "Variable 1"
	fmt.Println("first part of the string", string1, "Second part of the string")
	const constantValue = 100
	remainingValue := 50
	remValue := 50
	// Non re-assignable constantValue
	fmt.Println(constantValue)
	fmt.Println(remainingValue)
	fmt.Printf("The remValue id %v", remValue)
	var input int // int means whole numbers
	fmt.Scanf("Enter the inpu %d ", &input)
}
