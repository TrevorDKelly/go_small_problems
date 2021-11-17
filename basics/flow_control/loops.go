package main

import "fmt"

func main() {
  x := 0
  for {
    fmt.Println(x)
    x++
    if x >= 3 {
      break
    }
  }

  y := 0
  for y < 3{
    fmt.Println(y)
    y++
  }

  for x := 0; x < 3; x++ {
    fmt.Println(x)
  }

  for x := 0; x <= 100; x++ {
    if x % 2 == 0 || x % 3 == 0 {
      continue
    }
    fmt.Println(x)
  }
}
