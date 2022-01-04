/* Write a function that takes a positive integer n, sums all the cubed values from 1 to n, and returns that sum.

Assume that the input n will always be a positive integer. */

/*

  get int
  return int

  for each number from 1 to given. add its cubed value to a running total
  return the total

*/

package main

import "fmt"

func SumCubes(number int) int {
  total := 0
  for i := 1; i <= number; i++ {
    total += (i * i * i)
  }
  return total
}

func test(number, expected int) {
  result := SumCubes(number)
  fmt.Println(result == expected)
}

func main() {
  test(1, 1)
  test(2, 9)
  test(3, 36)
  test(123, 58155876)
}
