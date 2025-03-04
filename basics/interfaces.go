package main 

import "fmt"

type animal struct{
  name string
  age int
}
type human struct {
  deed string
}

//interfaces 
//Even if it can work without an interface , it provides more flexibility, polymorphism, etc
type speaks interface {
  speak() string
}

func (h human) speak() string {
  return "guarararara"
}
func (a animal) speak() string {
  return "meowwww....."
}

//polymorphism and flexibility
func makeSound(s speaks) {
  fmt.Println(s.speak())
}

func main() {
  var x speaks;

  x = human{}

  //Type assertions in Golang
  if _ , ok := x.(human) ; ok{
    fmt.Printf("A Human is now using the interface for speaking\n")
  } else {
    fmt.Printf("Something else I guess\n")
  }

  //Let's use Type Switches
  
  switch x.(type) {
  case animal:
    fmt.Printf("An animal is used\n")
  case human:
    fmt.Printf("A human is used\n")
  default:
    fmt.Printf("Default case runs\n")
  }
  ani := animal {
    name: "Cat",
    age: 12,
  }

  hoo := human {
    deed : "Die Live",
  }

  makeSound(ani)
  makeSound(hoo)
}
