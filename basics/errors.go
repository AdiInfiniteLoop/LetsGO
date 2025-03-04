package main 
//Error at the end is an interface having Error() as method behaivour
import (
  "fmt"
  "strconv"
  "errors"
)

//Custom Error Interface

type Divide struct {
  dividend int
  divisor int
}
type err interface {
  Error() string
}


func (de Divide) Error() string {
  return fmt.Sprintf("Cannot divide %d by %d\n",de.dividend, de.divisor,
)
}

func div(de Divide) (int, error){
  if   de.divisor == 0 {
    return -1, &de
  }
  ans :=  de.dividend / de.divisor
  return ans, nil 
}

func div2(de Divide) (float64, error) {
  if de.divisor == 0 {
  return -1.00, errors.New("Cannot divide by Zero\n")
  }
  return (float64)(de.dividend / de.divisor), nil
}
func main () {
  i, err := strconv.Atoi("100") //ascii to integer
  if err != nil {
    fmt.Println(err.Error())
    return
  }
  fmt.Println("No error %v", i)

  //Implementation of custom errors using Custom made interfaces
  de := Divide {
    dividend: 312,
    divisor: 0,
  }
  
  // We get the err here as the sender/ original de Divide that has the Error() implemented
   _, err = div(de)
  if err != nil {
    fmt.Printf(err.Error())
  } else {
  fmt.Printf("Let's go. Division Done by a Gopher\n")
  }


  //Implementation of custom errors using 'errors' package
    _, err = div2(de)
  if err != nil {
    fmt.Printf(err.Error())
  } else {
  fmt.Printf("Let's go. Division Done by a Gopher\n")
  }

  
}
