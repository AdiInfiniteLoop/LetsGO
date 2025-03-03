package main

import "fmt"
type Vehicle struct {
  Name string
  Model string
}

//Struct Embedding
type Car struct {
  Vehicle
  Owner string
}

//struct methods
func (c Car)concatnames() string {
  return c.Name + c.Model + c.Owner
}

func main() {

  mycar := Car{
    Vehicle: Vehicle{Name: "MOtootot", Model: "XXD23"},
    Owner: "ADitya",
  }
  fmt.Printf("%s\n", mycar.Name)

  //Anonymous Structs (Used in HTTP Handler)
  myanimal := struct {
    Name string
    Age int
  } {
    Name: "Tiger Panda",
    Age: 21,
  }

  fmt.Printf("%d\n", myanimal.Age)

   con := mycar.concatnames()
   fmt.Printf("%s\n", con)
}
