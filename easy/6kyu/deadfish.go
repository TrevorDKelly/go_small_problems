/* Write a simple parser that will parse and run Deadfish.

Deadfish has 4 commands, each 1 character long:

i increments the value (initially 0)
d decrements the value
s squares the value
o outputs the value into the return array
Invalid characters should be ignored. */

/*

given an input string run the series of actions they represent. return the array the actions create

get string
return array of ints

build struc with methods for each action properties of value and array

init deadfish
iterate thru string
  call val associated with it
return deadfish collection

*/

package main

import "fmt"

type Deadfish struct {
  value int
  collection []int
}

func (fish *Deadfish) pass(action rune) {
  switch action {
  case 'i':
    fish.value += 1
  case 'd':
    fish.value -= 1
  case 's':
    fish.value *= fish.value
  case 'o':
    fish.collection = append(fish.collection, fish.value)
  }
}

func Parse(str string) []int {
  fish := Deadfish{0, []int{}}
  for _, char := range str {
    fish.pass(char)
  }
  return fish.collection
}

func main() {
  fmt.Println(Parse("ssdizbsnsdkkincuybidds"))
}
