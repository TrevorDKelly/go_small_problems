/*  Calculate how many years ago the father was twice as old as his son (or in how many years he will be twice as old).*/

/*

  get int, int
  return int

  with the current ages, find how many years ago the oler was twice the age of the younger

  get age difference
  difference from sons age

  dad - son
  son - diff
  convert to positive
    if negative
      * -1

  return

*/

package main

import "fmt"

func TwiceAsOld(father, son int) int {
  diff := father - son
  ans := son - diff
  if (ans < 0) {
    ans *= -1
  }
  return ans
}

func test(dad, son, expected int) {
  result := TwiceAsOld(dad, son)
  fmt.Println(result == expected)
}

func main() {
  test(36, 7, 22)
  test(55, 30, 5)
  test(42, 21, 0)
}
