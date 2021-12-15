package main

import "fmt"

func main() {
  myArr := [5]int{0, 1, 2, 3, 4}

  // var s []string = myArr[2:4]
  var s []int = myArr[2:4]
  fmt.Println(s)

  s2 := myArr[3:5]
  fmt.Println(s2)

  // s3 := myArr[3:6]
  // fmt.Println(s3)

  fullSlice := myArr[:]

  s4 := fullSlice[3:5]
  fmt.Println(s4)

  // s5 := fullSlice[4:6]
  // fmt.Println(s5)

  s6 := s4[2:]
  fmt.Println(s6)

  s7 := myArr[:3]
  s8 := s7[:]
  fmt.Println(cap(s8))

  var s9 []int = []int{0, 1, 2, 3}
  fmt.Println(cap(s9))

  s10 := []int{}
  fmt.Println(s10)
  fmt.Println(cap(s10))

  s10 = append(s10, 0)
  fmt.Println(s10)

  s10 = append(s10, 1, 2)
  fmt.Println(s10)

  // s10 = append(s10, s9)
  // fmt.Println(s10)

  s10 = append(s10)
  fmt.Println(s10)

  s11 := make([]int, 5)
  fmt.Println(s11)
  fmt.Println(cap(s11))

  // s12 := make(s10)
  // fmt.Println(s12)


  var myArr2 [5]int
  fmt.Println(myArr2)

  s13 := myArr2[:4]
  myArr2[2] = 3
  fmt.Println(s13)
  fmt.Printf("s13 cap is %v\n", cap(s13))
  s13 = append(s13, 5)
  fmt.Printf("s13 cap is %v\n", cap(s13))
  myArr2[3] = 4
  fmt.Println(s13)
  // same underlying array after appending
  s13 = append(s13, 6)
  fmt.Printf("s13 cap is %v\n", cap(s13))
  myArr2[1] = 2
  fmt.Println(myArr2)
  fmt.Println(s13)
  // if append surpasses the underlying array length a new one is created
  s14 := s13[:]
  fmt.Printf("s14 cap is %v\n", cap(s14))
  s13 = append(s13, 7, 8, 9, 10, 11)
  s13[0] = 1
  fmt.Printf("s13 cap is %v\n", cap(s13))
  fmt.Printf("s14 cap is %v\n", cap(s14))
  fmt.Println(s14)

  //copy
  slice := []int{1,2,3,4,5}
  first3 := make([]int, 3)

  copy(first3, slice[:3])
  fmt.Println(first3)
  copy(first3, slice[1:])
  fmt.Println(first3)
  copy(first3[1:], slice[:3])
  fmt.Println(first3)
}
