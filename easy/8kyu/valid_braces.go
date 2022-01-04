/* Write a function that takes a string of braces, and determines if the order of the braces is valid. It should return true if the string is valid, and false if it's invalid. */

/*

  get string
  return boolean

  given a string of opening and closing parenth, brackets, and sq brackets
  check if they are in an order that makes a valid set of containers

  for every open
    make sure everything inside closes before it closes

  when we come across an open
    create object that is open and has a brace type
  when a close is found
    make sure every other open after it is closed

  iterate
    if open is found
      make new open
      add to list
    if close is found
      if last was it's open
        remove both
      if not return false

  init open braces as empty slice
  iterate thru string - runes
    make string
    if an open brace
      add to slice
    if close
      if last brace in slice is its match
        remove last slice element
      else return false
  slice empty ?

*/

package main

import "fmt"

func ValidBraces(str string) bool {
  var openBraces []string

  for _, runeLetter := range str {
    letter := string(runeLetter)

    if letter == "(" || letter == "[" || letter == "{" {
      openBraces = append(openBraces, letter);
    } else {
      length := len(openBraces)
      if length == 0 {
        return false
      }

      last := openBraces[length - 1]
      if (last == "(" && letter != ")") ||
         (last == "[" && letter != "]") ||
         (last == "{" && letter != "}") {
        return false
      } else {
        openBraces = openBraces[:length - 1]
      }
    }
  }

  return len(openBraces) == 0
}

func test(str string, expected bool) {
  result := ValidBraces(str)
  fmt.Println(result == expected)
}

func main() {
  test("()", true)
  test("[]", true)
  test("{}", true)
  test("{()()([])}", true)
  test("([{}])", true)
  test("(}", false)
  test("()}", false)
  test("({)}", false)
}
