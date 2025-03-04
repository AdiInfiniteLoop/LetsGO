package main 

import "fmt"
//Variadic functions
//Each and every element in the slice should pass 
//...Type indicates that the function can accept a variable number of arguments of type Type
func calcSum(array ...int) (int, error){
  sum := 0
  for i := 0; i < len(array); i++ {
    sum += array[i];
  }
  return sum, nil
}


func main() {
  //Arrays
  //var arr [3]int;
  //arr2 := [5]int {1,23,3,4,3}
  
  //Slices
  //myslice := arr2[1:3] //not include 3rd
  //allelements := arr2[:]


  //Make Slices
  //myslice := make([]int, 5)
  myslice := []int{1,2,3,4,5}
  //myslice := make([]int, 5, 10) where 10 is the max length or length of underlying array 
  //Spread operator to pass the whole slice
  
  //Append is a variadic function so any amount of elements can be appended
  myslice = append(myslice, 6,7,8)
  ans, _:= calcSum(myslice...)
  fmt.Println(ans)

  //SLice of Slices 2 x 2
  //slc := make([][] int, 3); [[][][]]

  for idx,  it := range myslice {
    fmt.Println(idx, it)
  }

}
