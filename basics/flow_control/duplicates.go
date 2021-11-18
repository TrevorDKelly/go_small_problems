package main

import "fmt"

func main() {
  fullArr := [10]int{3,54,3,2,6,90,8,42, 4, 2}

  for i, element := range fullArr {
    for _, element2 := range fullArr[i + 1:] {
      if element == element2 {
        fmt.Println(element)
      }
    }
  }
}
