package main

import "fmt"

func main() {
  var map1 map[int]string
  fmt.Println(map1)
  map2 := map[int]string{ 1: "this" }

  fmt.Println(map2)

  val, ok := map2[1]
  fmt.Println(val, ok)
}
