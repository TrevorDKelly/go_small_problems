/*  Create a function that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers.*/

/*



*/

package main

import "fmt"

func EvenOrOdd(number int) string {
  if number % 2 == 0 {
    return "Even"
  }
  return "Odd"
}


func test(number int, expected string) {
  result := EvenOrOdd(number)
  fmt.Println(result == expected)
}

func main() {
  test(2, "Even")
  test(-2, "Even")
  test(0, "Even")
  test(1, "Odd")
  test(-1, "Odd")
}
