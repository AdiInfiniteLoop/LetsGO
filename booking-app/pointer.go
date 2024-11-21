package main


import "fmt"

func main() {

	fmt.Println("Let's point to Pointers")


	//var ptr *int;	//initially value is "nil" when not assigned

	mynumber := 23
	// var ptr *int = &mynumber
	var ptr  = &mynumber
	fmt.Printf("the value is %v \n", *ptr)
	*ptr = *ptr *13
	fmt.Printf("the new value is %v \n", *ptr)
	



}