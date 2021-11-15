package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Printf("Enter your birth year: ")
  scanner.Scan()

  input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

  fmt.Printf("You'll be %d years old at the end on the year", 2020 - input)
}
