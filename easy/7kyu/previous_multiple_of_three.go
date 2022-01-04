/*Given a positive integer n: 0 < n < 1e6, remove the last digit until you're left with a number that is a multiple of three.

Return n if the input is already a multiple of three, and return null/nil/None/-1 if no such number exists.  */

/*

  get int
  return int/nil

  given a number check if it is divisible by 3.
  if not remove the last digit and check again
  return the first number found that is divisible by 3 or nil

  while num > 0
    if divisible by 3 return
    else set num to num / 10
  return nil


*/

package main

import "fmt"

func PrevMultOfThree(number int) (result interface{}) {
  for number > 0 {
    if number % 3 == 0 {
      result = number
      return
    } else {
      number /= 10
    }
  }
  return
}

func test(num int, expected interface{}) {
  result := PrevMultOfThree(num)
  fmt.Println(result == expected)
}

func main() {
  test(1, nil)
  test(25, nil)
  test(1244, 12)
  test(952406, 9)
}
