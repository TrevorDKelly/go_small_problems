/* Consider an array/list of sheep where some sheep may be missing from their place. We need a function that counts the number of sheep present in the array (true means present). */

/*

  get array of bool
  return int

  init count
  loop over array
    if true add 1
  return count
*/

package main

import "fmt"

func countSheeps(sheeps []bool) int {
  count := 0

  for _, there := range sheeps {
    if there {
      count += 1
    }
  }

  return count
}

func test(sheeps []bool, expected int) {
  result := countSheeps(sheeps)
  fmt.Println(result == expected)
}

func main() {
  sheeps := []bool{true,  true,  true,  false,
  true,  true,  true,  true ,
  true,  false, true,  false,
  true,  false, false, true ,
  true,  true,  true,  true ,
  false, false, true,  true}

  test( sheeps, 17)
}
