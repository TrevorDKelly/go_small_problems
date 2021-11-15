package main

import "fmt"

func main() {
  zero := 0
  one := 1
  str := "this"
  eStr := ""

  //zeroTest := true && zero
  //numTest := true && one
  //strTest := true && str
  //eStrTest := true && eStr

  // fmt.Printf("zero %v", zeroTest)
  // fmt.Printf("number %v", numTest)
  // fmt.Printf("stirng %v", strTest)
  // fmt.Printf("empty string %v", eStrTest)
  // ERRORS RAISED
  //   && not defined on non-boolean types

  // test short circuting

  // fmt.Println(true || zero)
  // does not work

  fmt.Println(true || !!(one))
}
