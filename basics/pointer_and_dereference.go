package main

import "fmt"

func iterate(num *int) {
  *num += 1
}

func main() {
  x := 1
  y := &x
  z := x
  fmt.Printf("%T : %T\n", x, y)
  fmt.Println(*&x)

  *&x = 2
  fmt.Println(x, z)

  iterate(&x)
  fmt.Println(x)

  // a := &"str"
  // fmt.Println(a)

  arr := [3]int{1,2,3}

  refArr := &arr

  fmt.Println(refArr)
  fmt.Println(refArr[0])

  refArr[0] = 100

  fmt.Println(arr)

  slice := arr[:]
  refSlice := &slice

  fmt.Println(refSlice)
  // refSlice[1] = 123
  // *refSlice[1] = 123
  (*refSlice)[1] = 123
  fmt.Println(arr)

  refEl := &slice[2]
  fmt.Println(refEl)
  *refEl = 789
  fmt.Println(arr)

  myMap := map[int]string{
    1: "one",
    2: "two",
  }

  refMap := &myMap

  fmt.Println(refMap)

  // refMap[1] = "ONE"
  // *refMap[1] = "ONE"
  // fmt.Println(refMap)

  fmt.Printf("%T : %T", refMap, *refMap)
}
