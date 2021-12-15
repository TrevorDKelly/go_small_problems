/* Write a function called repeatStr which repeats the given string string exactly n times. */

/*
  get int string
  return string

  return a new string that is the given string repeated int times

  init new string
  loop int times
    add string to new string
  return new string

*/

package main

import "fmt"

func RepeatStr(times int, str string) string {
  var newString string

  for i := 0; i < times; i++ {
    newString += str
  }

  return newString
}

func test(times int, str, expected string) {
  result := RepeatStr(times, str)
  fmt.Println(result == expected)
}

func main() {
  test(6, "I", "IIIIII")
  test(3, "Hello", "HelloHelloHello")
}
