package main

import (
  "fmt"
)

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
//Channels are safe queue for go routines
//as there are no returns types for goroutines in the classical sense thus
//we use Channels to store and share data


  /*`Channel operator is blocked until another go routine  is ready to execute`
    ->Think of a channel like a two-way door between rooms. 
    When you try to send something through the door (using <-), 
    you need someone on the other side to receive it. 
    If nobody's there, you can't just push the thing through - 
    you have to wait until someone opens the door from the other side.
  */
  channel := make(chan bool)

  go func(limit int) {
    for i:= 1; i <= limit; i++{
      if i % 2 == 0 {
        channel <- true 
        continue
      }
    channel <- false
    }
  }(4)
//  if i put limit as 2 and do 3 print statements then channel cannot store values as there ios 
//Yes, exactly! The main function is waiting for the concurrent goroutine 
//(the func() that runs inside go func() { ... }) to update the channel.
  val := <-channel
  fmt.Println(val, "-") 
  val = <-channel
  fmt.Println(val, "-")
 val = <-channel
  fmt.Println(val, "-")
  
  //Buffered channels only blocks when buffer is full
  //close the channel (Always from the sender side)
  close(channel)

  //**val, ok := <-channel  //to check if channel still available
  

  chints := make(chan int)
   go fibonacci(8, chints);

  for i := range chints {
    fmt.Println(i)
  }
    fmt.Println("Channel is closed, all values received.")
/*   select {
      case email, ok:= <- chEmails:
    // This block is executed when ch1 has a value to receive.       
        if !ok {
          return 
        }
      case msg, ok <- chMsgs:
        if !ok {
          ....
        }
      default: 

   } */
}
