/* In this Kata, we will check if a string contains consecutive letters as they appear in the English alphabet and if each letter occurs only once. */

/*

  get string
  return boolean

  check that there are no duplicate letters and that there are no gaps between letters when alphabetically listed

  import sort
  sort
  iterate runes
    subtract 1 from 2
    if its not -1 return false

  import sort strings

  split string
  sort
  get first
  iterate second to end
    if second - first is not 1
      return false
    first = second
  return true

  init min at 122
  init total at 0
  iterate thru runes
    if less than min
      min = rune
    add val to total
  multiply min by length of string
  subtract from total
  formula for what should match
  add length
*/

package main

import (
  "fmt"
)

func Solve(str string) bool {
  min := 122
  max := 97
  total := 0

  for _, val := range str {
    intVal := int(val)
    if intVal < min {
      min = intVal
    }
    if intVal > max {
      max = intVal
    }
    total += intVal
  }

  correctSum := (max - min + 1) * (min + max) / 2
  return total == correctSum
}

func test(str string, expected bool) {
  result := Solve(str)
  fmt.Println(result == expected)
}

func main() {
  test("a", true)
  test("ab", true)
  test("ac", false)
  test("abd", false)
  test("dabc", true)
  test("abbc", false)
}
