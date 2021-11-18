package main

import "fmt"

func main() {
  arr := [3]int{1, 2, 3}
  arr2 := arr
  slice := arr[1:]
  slice2 := arr[:2]

  //slice[0] = 10
  arr[1] = 10
  fmt.Println(arr, arr2, slice, slice2)
}
