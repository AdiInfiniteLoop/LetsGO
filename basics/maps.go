package main

import "fmt" 

func main() {

  age := make(map[string]int)
  age["adi"] = 21;
  fmt.Println(age["adi"])

  ages := map[string] int {
    "ADi": 21,
  }
  delete(age, "adi")
  fmt.Println(ages["ADi"])

  //ok is not error it is a bool val
  ele, ok := ages["Ado"]
  if ok != true {
    fmt.Println("Error occured. No such key found")
  } else {
    fmt.Println(ele)
  }
} 
