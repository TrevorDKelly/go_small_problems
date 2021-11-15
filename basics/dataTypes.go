package main

import "fmt"

func main() {
  var str string
  var num8 uint8
  var num64 uint64
  var myInt int
  var fl float64
  var bolean bool
  var byt byte

  // testing default values
  fmt.Println("Defalt string: ",  str)
  fmt.Println("Defalt uint8: ",  num8)
  fmt.Println("Defalt uint64: ",  num64)
  fmt.Println("Defalt int: ",  myInt)
  fmt.Println("Defalt float64: ",  fl)
  fmt.Println("Defalt bool: ",  bolean)
  fmt.Println("Defalt byte: ",  byt)

  // converting numbers

  // fmt.Println(num8 + num64) ERROR
  fmt.Println(uint64(num8) + num64)

  var neg64 int64 = -42
  fmt.Println(neg64 + int64(num64))

  // comparing strings
  a := "a"
  b := "b"
  A := "A"
  var s string

  fmt.Println(a < b)
  fmt.Println(a < A)
  fmt.Println(A < b)
  fmt.Println(A < s)
  fmt.Println(A > s)
  fmt.Println("" == "")

}
