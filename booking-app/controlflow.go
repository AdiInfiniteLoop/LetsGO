package main

import "fmt"

func main() {
	fmt.Println("Welcome to Control Flow")

	age := 21
	if age > 18 {
		fmt.Println("Good adult")
	}
	else if age == 18 {
		fmt.Println("Yolo")
	}
	else {
		fmt.Println("Bad Child")
	}

	/* We cannot use 
		if age > 18
		{
		}

		as { needs to be in the same line as if statement
	*/

	if num := 10; num < 10 {
		fmt.Println("Less than 10")
	}
	else {
		fmt.Println("Good enough")
	}
}

