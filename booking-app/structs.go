package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")

	//No inheritance in GO lang
	//No super, parent keyword
	luffy := User("name", "abc@mgmail.com", "124", true, 21)
	fmt.Println(luffy)
	fmt.Println(luffy.Name)
	fmt.Printf("The details are %+v \n", luffy) //use +v as it gives more details

}

//Make it capital as User is public
type User struct {
	Name string
	Email string
	Password string
	Status bool
	Age int64
}