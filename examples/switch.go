package main

import(
  "fmt"
)

func main(){
  var i int
  fmt.Print("Pick of Number:")
  fmt.Scan(&i)

  switch i{
  case 1:
    fmt.Println("You picked one!")
  case 2:
    fmt.Println("You picked two!")
  case 3:
    fmt.Println("You picked Three!")
  }
}
