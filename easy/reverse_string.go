/* Complete the solution so that it reverses the string passed into it. */

/*
  get string
  return string

  return string should have all letters reversed

  iterate thru letters
    for each add collected letters to the end of it

  init new string
  iterate length of given string
    get letter at i
    reassign new string to letter + new string
  return new string
*/

package main

import "fmt"

func Solution(word string) string {
  var reversed string

  for _, letter := range word {
    reversed = string(letter) + reversed
  }

  return reversed
}

func test(word, reverse string) {
  result := Solution(word)
  fmt.Println(result == reverse)
}

func main() {
  test("the", "eht")
  test("a", "a")
  test("", "")
}
