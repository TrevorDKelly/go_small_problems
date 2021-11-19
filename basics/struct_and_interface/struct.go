package main

import "fmt"

type Dog struct {
  name string
  age int
}

func main() {
  var dog1 Dog = Dog{"spot", 3}
  fmt.Println(dog1)

  dog2 := Dog{"Seamus", 6}
  fmt.Println(dog2)

  dog3 := Dog{age: 3}
  fmt.Println(dog3)

  dog4 := Dog{}
  fmt.Println(dog4)

  // dog5 := Dog{"Jim"}
}
