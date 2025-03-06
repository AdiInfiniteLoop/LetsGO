package main 

import (
  "fmt"
  "mystrings"
) 

//Here mystrings and conc are two different modules

//to export a function we must capitalize it else not
//eg. Hello() can be exported but not hello()
func main() {
  fmt.Println("Build production executable")
  //PACKAGE_NAME.Hello()
   fmt.Println(mystrings.Reverse("ddddaaaa"));
}
