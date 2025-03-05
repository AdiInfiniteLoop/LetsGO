package main 

import "fmt"

//Currying Functions: When we pass a function as an input and returns a new function
func selfAdd(a int) func(int) int {
  x := 80
  return func(a int) int {
    return a + a + x  //here the function is a closure also as it accesses x that is outside it's scope. so it can mutate x also.
  }
}

func main() {
  ele := selfAdd(10)
  fmt.Println(ele(10))
}
