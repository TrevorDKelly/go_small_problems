/*  You are given two interior angles (in degrees) of a triangle.

Write a function to return the 3rd.
*/

/*

  get int, int
  return int

  a triangle adds up to 180 degrees with internal angeles. find the third angle given the other two

  add two ints together then subtract from 180

*/

package main

import "fmt"

func OtherAngle(first, second int) int {
  total := first + second
  return 180 - total
}

func test(one, two, expected int) {
  result := OtherAngle(one, two)
  fmt.Println(result == expected)
}

func main() {
  test(30,60,90)
  test(60,60,60)
  test(43,78,59)
}
