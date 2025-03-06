package main

import (
  "fmt"
  "time"
)
func main() {
  //go keyword for doing goroutines concurrently
go func () {
    fmt.Println("Concurrently happening") 
  }()
    time.Sleep( time.Millisecond * 249)
  fmt.Println("outside..")
}
