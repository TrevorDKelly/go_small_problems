package main

import "fmt"

func main() {
  boolean := true
  str := "string"
  num := 42

  fmt.Printf("boolean boolean: %t \n", boolean)
  fmt.Printf("boolean value: %v \n", boolean)
  fmt.Printf("boolean string: %t \n", str)
  fmt.Printf("string type: %T \n", str)
  fmt.Printf("num value: %v \n", num)

  fmt.Printf("stirng with space: %9v\n", str)

  fmt.Printf("Two interpolations one variable: %T, %v\n", str)
}
