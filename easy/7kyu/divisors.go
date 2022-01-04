/* Count the number of divisors of a positive integer n. */

/*

get int
return int

for a number count how many divisors it has

special case for 1
  returns 1
all other start at 2
get square root
  find if whole number
    if it is add 1 to count
iterate 2 to sqr root -1
  if modulo num and i is 0 add two to count
return count

import math
if number is 1 return 1
init count at 2
get sqare root
if sqr is whole number add 1 to count
change sqr to int
iterate 2 to intSqr - 1
  if num % 1 is 0
    add 2 to count
return count

*/

package main

import (
  "fmt"
  "math"
)

func Divisors(number int) int {
  if number == 1 {
    return 1
  }

  count := 2
  square := math.Sqrt(float64(number))
  intSquare := int(square)
  if square == float64(intSquare) {
    count += 1
    intSquare -= 1
  }

  for i := 2; i <= intSquare; i++ {
    if number % i == 0  {
      count += 2
    }
  }

  return count
}

func test(number, expected int) {
  result := Divisors(number)
  fmt.Println(result == expected)
}

func main() {
  test(1,1)
  test(10, 4)
  test(11, 2)
  test(54, 8)
}
