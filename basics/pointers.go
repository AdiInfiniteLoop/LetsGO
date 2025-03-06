package main 

import "fmt"

type Car struct  {
  name string
  model string
}

//Pass the referenced struct obj
func (c *Car) show() {
  fmt.Println(c.name)
  c.name = "dddderre12"
}
func main() {

  var ptr *string;  //string is the type it is referring
  stringname := "Adi"
  ptr = &stringname
  fmt.Println(*ptr)
  fmt.Println(&stringname)

  //Check nil pointers as it will cause panic when deferencing is done
  
  c := Car{
    name: "XSSE1",
    model: "24233ss",
  }
  //Need not pass with '&'
  c.show()
  fmt.Println(c.name)
}
