/* In this Kata, you will be given a number n (n > 0) and your task will be to return the smallest square number N (N > 0) such that n + N is also a perfect square. If there is no answer, return -1 (nil in Clojure, Nothing in Haskell, None in Rust).

More examples in test cases.

Good luck! */

/*

  get int
  return int

  find the square number for which given int can be added to it to produce another square number. return -1 if no answer

  Break point
  find the square at wich the difference between two squares jumps higher than the given number

  init last square at 1
  iterate 1 -
    get the square of i
    if square - last square > n
      return - 1
    add N to square
    get square root
    check if whole number
      is to int to float == float
      if yes
        return square as int
    save square as last square

*/

package main

import (
  "fmt"
  "math"
)

func Solve(n int) int {
  lastSquare := 1
  for i := 1;; i++ {
    square := i * i
    if square - lastSquare > n {
      return -1
    }

    root := math.Sqrt(float64(square + n))
    if float64(int64(root)) == root {
      return square
    }
    lastSquare = square
  }
}

func test(num, expected int) {
  result := Solve(num)
  fmt.Println(result == expected)
}

func main() {
  test(1, -1)
  test(3, 1)
  test(5, 4)
  test(9, 16)
  test(10, -1)
  test(88901, 5428900)
}
