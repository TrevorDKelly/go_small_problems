package main

import "fmt"

func main() {
  var myArr [3]int

  fmt.Println(myArr)
  fmt.Println(myArr[0])

  myArr[0] = 5

  fmt.Println(myArr[0])

  // myArr[5] = 1
  // fmt.Println(myArr[4])

  // var arr2 [3]string{"this", "that", "the other"}
  arr2 := [3]string{"this", "that", "the other"}
  fmt.Println(arr2)

  var nested [2][2]int
  fmt.Println(nested)

  // nested[1] := [2]int

  // var inner [4]int
  // nested[1] = inner

  // arr := make([2]int, 2)
  // fmt.Println(arr)
}
