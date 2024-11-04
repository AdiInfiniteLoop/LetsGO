package main

import "fmt"

func main() {
	fmt.Println("Welcome users to our Booking App")
	// := is called Short Variable Declaration(SVD). But I call it gosign
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
	fmt.Scanf("Enter the input %d ", &input)
	/*var x,y,z int = 23; used for declaring multiple variables at once
	a := 23; cannot explicitly define variables data type
	Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
	*/

	// Syntax for slices in go
	// slice := make([]string, 0, 40) // datatype, fill with 0, capacity
	// Since := both declares and intializes value
	// arr2 := []string{} // empty slice
	// arr3 := []string{"SliceInitial Val1", "SliceInitial Val2"}

	// Loops: There is only for loop in go
	// 1.Infinite Loop
	// for {
	// fmt.Println("Aditya")
	// }
	// 2.For each Loops
	slice := []string{"1", "2", "3", "4", "5"}
	for index, element := range slice {
		fmt.Print(element)
		fmt.Println(index)
	}
	// Doesn't include 10 and starts from 0
	for index := range 10 {
		fmt.Print(index)
	}

	// infinite loop causes unreachable code as it takes 100% CPU
	// for {
	// any operation
	// }
	for _, num := range 100 {
		fmt.Printf(num)
	}
}
