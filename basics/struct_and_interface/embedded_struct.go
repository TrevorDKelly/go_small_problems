package main

import "fmt"

type Name struct {
  first string
  last string
}

type Person struct {
  name *Name
  age int
}

type Person2 struct {
  *Name
  age int
}

type Overlap struct {
  first bool
  *Name
}

func main() {
  var joe Person = Person{&Name{"Joe", "Shmo"}, 98}
  fmt.Println(joe)

  bethName := Name{"Beth", "Smith"}
  beth := Person{&bethName, 49}
  fmt.Println(beth.name.first)
  // fmt.Println(beth.first)

  beth2 := Person2{&bethName, 74}
  fmt.Println(beth2.first)

  overBeth := Overlap{true, &bethName}
  fmt.Println(overBeth)
  fmt.Println(overBeth.first)
  fmt.Println(overBeth.Name.first)
}
