/* Given a Divisor and a Bound , Find the largest integer N , Such That ,

Conditions :
N is divisible by divisor

N is less than or equal to bound

N is greater than 0.
*/


/*

  get int, int
  return int


  find max multiple of first int that is less than second int

  divide second by first
  multiply result times first

*/

package main

import "fmt"

func MaxMultiple(divisor, bound int) int {
  maxMultiplier := bound / divisor

  return divisor * maxMultiplier
}

func test(divisor, bound, expected int) {
  result := MaxMultiple(divisor, bound)
  fmt.Println(result == expected)
}

func main() {
  test(2,7,6)
  test(3,10,9)
  test(7,17,14)
  test(10,50,50)
  test(37,200,185)
  test(7,100,98)
}
