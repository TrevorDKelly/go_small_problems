/* The two oldest ages function/method needs to be completed. It should take an array of numbers as its argument and return the two highest numbers within the array. The returned value should be an array in the format [second oldest age, oldest age].

The order of the numbers passed in could be any order. The array will always include at least 2 items. If there are two or more oldest age, then return both of them in array format.

 */

/*

  get slice
  return array

  get the two largest numbers in the slice and return them as an array


  iterate thru slice. if number is bigger that the largest, set smaller to largest and largest to number. if number is bigger than smaller, set smaller to num

  init oldest, secondOldest to 0
  iterate thru slice
    if bigger than oldest
      second oldest = oldest
      oldest = num
    if bigger than second
      second = num

  return new array
*/

package main

import "fmt"

func TwoOldestAges(ages []int) [2]int {
  oldest := 0
  secondOldest := 0

  for _, age := range ages {
    if age > oldest {
      secondOldest = oldest
      oldest = age
    } else if age > secondOldest {
      secondOldest = age
    }
  }
  return [2]int{secondOldest, oldest}
}

func test(ages []int, expected [2]int) {
  result := TwoOldestAges(ages)
  fmt.Println(result == expected)
}

func main() {
  test([]int{6,5,83,5,3,18}, [2]int{18,83})
  test([]int{1,5,87,45,8,8}, [2]int{45,87})
}
