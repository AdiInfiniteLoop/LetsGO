package main 

import "fmt"

//Currying Functions: When we pass a function as an input and returns a new function
func selfAdd(a int) func(int) int {
  return func(a int) int {
    return a + a
  }
}

func main() {
  ele := selfAdd(10)
  fmt.Println(ele(10))
}
