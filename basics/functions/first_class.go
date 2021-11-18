package main

import "fmt"

func add1(val int) int {
  return val + 1
}

func main() {
  add2 := func (val int) int {
    return val + 2
  }

  adder := func(val int, one, two func(int) int) int {
    return one(val) + two(val)
  }

  fmt.Println(adder(1, add1, add2))

  returnAFunc := func(val int) func(int) string {
    return func(val2 int) string {
      if val > val2 {
        return "not as big"
      } else {
        return "bigger or equal"
      }
    }
  }

  comparedTo3 := returnAFunc(3)

  fmt.Println(comparedTo3(2))
  fmt.Println(comparedTo3(4))
}
