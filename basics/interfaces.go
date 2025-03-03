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
