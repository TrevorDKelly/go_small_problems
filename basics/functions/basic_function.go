package main

import "fmt"

func sayName(name string, age int) {
  fmt.Printf("my name is %v and i am %v.\n", name, age)
}

func oneReturn(val int) int {
  val++
  return val
}

func multipleReturns(val int) (int, int) {
  return val + 1, val - 1
}

func namedReturns(val int) (less, more int) {
  less = 3
  return
}

func typesInARow(val3 string, val1, val2 int) {
  fmt.Println(val1, val2, val3)
}

func add1(val int) {
  val++
}

func whenToDefer(val int) int{
  defer fmt.Println(val)

  add1(val)
  fmt.Println(val)
  return val
}

func main() {
  // sayName("Trevor", 31)
  // fmt.Println(oneReturn(1))
  // fmt.Println(multipleReturns(2))
  // fmt.Println(namedReturns(2))
  // typesInARow("this",1,2)

  whenToDefer(1)

  // after()
}

func after() {
  fmt.Println("does this work")
  // Yes
}
