package main

import "fmt"

func byValue(str string) {
  str = "changed by value"
}

func byReference(str *string) {
  *str = "changed by reference"
}

func returnValue(str *string) string {
  return *str
}

func returnReference(str *string) *string {
  return str
}

//arrays

func arrByValue(arr [3]int) {
  arr[1] = 100
}

func arrByReference(arr *[3]int) {
  arr[1] = 200
}

func arrReturnValue(arr *[3]int) [3]int {
  return *arr
}

func arrReturnReference(arr *[3]int) *[3]int {
  return arr
}

//slices

func sliceByValue(slice []int) {
  slice[1] = 100
}

func sliceByReference(slice *[]int) {
  (*slice)[1] = 200
}

func sliceReturnValue(slice *[]int) []int {
  return *slice
}

func sliceReturnReference(slice *[]int) *[]int {
  return slice
}

//maps

func mapByValue(myMap map[int]string) {
  myMap[1] = "ONE"
}

func mapByReference(myMap *map[int]string) {
  (*myMap)[1] = "TWO"
}

func mapReturnValue(myMap *map[int]string) map[int]string {
  return *myMap
}

func mapReturnReference(myMap *map[int]string) *map[int]string {
  return myMap
}


func main() {
  str := "My String"

  byValue(str)
  fmt.Println(str)

  byReference(&str)
  fmt.Println(str)

  str2 := returnValue(&str)
  *&str2 = "can change when value returned"
  fmt.Println(str, str2)

  str3 := returnReference(&str)
  *str3 = "returned reference is reference to same object"
  fmt.Println(str)

  //arrays
  arr := [3]int{1,2,3}

  arrByValue(arr)
  fmt.Println(arr)

  arrByReference(&arr)
  fmt.Println(arr)

  arr2 := arrReturnValue(&arr)
  arr2[1] = 300
  fmt.Println(arr)

  arr3 := arrReturnReference(&arr)
  arr3[1] = 400
  fmt.Println(arr)

  //slices
  slice := []int{1,2,3}

  sliceByValue(slice)
  fmt.Println(slice)

  sliceByReference(&slice)
  fmt.Println(slice)

  slice2 := sliceReturnValue(&slice)
  slice2[1] = 300
  fmt.Println(slice)

  slice3 := sliceReturnReference(&slice)
  (*slice3)[1] = 400
  fmt.Println(slice)

  //maps
  myMap := map[int]string{
    1:"one",
    2:"two",
    3:"three",
    4:"four",
  }

  mapByValue(myMap)
  fmt.Println(myMap)

  mapByReference(&myMap)
  fmt.Println(myMap)

  myMap2 := mapReturnValue(&myMap)
  myMap2[1] = "THREE"
  fmt.Println(myMap)

  myMap3 := mapReturnReference(&myMap)
  (*myMap3)[1] = "FOUR"
  fmt.Println(myMap)

}
