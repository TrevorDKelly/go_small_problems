/* Return the number (count) of vowels in the given string.

We will consider a, e, i, o, u as vowels for this Kata (but not y).

The input string will only consist of lower case letters and/or spaces. */

/*

  get string
  return int

  return the number of vowels in the given string

  hash table of vowels
  itererate thru letters
    if in hash table add 1 to count

  create vowels hash table
    rune: true

  init count at 0
  iterate thru string
    if vowels[letter]
      add 1 to count
  return count

*/

package main

import "fmt"

var VOWELS map[rune]bool = map[rune]bool{
  'a': true,
  'e': true,
  'i': true,
  'o': true,
  'u': true,
}

func GetCount(str string) (count int) {
  for _, char := range str {
    if VOWELS[char] {
      count++
    }
  }
  return count
}

func test(str string, expected int) {
  result := GetCount(str)
  fmt.Println(result == expected)
}

func main() {
  test("aaa", 3)
  test("abcde", 2)
  test("aeiouy", 5)
  test("", 0)
  test("mt my", 0)
}
