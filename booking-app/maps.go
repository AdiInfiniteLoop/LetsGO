package main

import "fmt"

func main() {
	//HashTables

	fmt.Println("Welcome to Hash Tables in Go")

	hashmap := make(map[string]int) //map[key datatype]value datatype
	hashmap["typescript"] = 10
	hashmap["cpp"] = 100
	hashmap["go"] = 1000

	fmt.Println("The map is", hashmap)
	fmt.Println("The map is", hashmap["typescript"])

	delete(hashmap, "typescript") //delete the typescript

	//SImilar to for-each loop
	for key, value := range hashmap {
		fmt.Println(key, " - ", value)
	}
	

}