package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  /*
  for {
    scanner.Scan()
    ans, _ := strconv.ParseInt(scanner.Text(), 10, 64)

    switch ans {
    case 0:
      fmt.Println("it's zero")
    case 1:
      fmt.Println("it's one")
    default:
      fmt.Println("other")
    }
  }
  */

  for {
    scanner.Scan()
    ans := scanner.Text()
    num, _ := strconv.ParseInt(ans, 10, 64)

    switch {
    case num != 0 && ans != "0":
      fmt.Println("its a number")
    case ans == "0":
      fmt.Println("its zero")
    default:
      fmt.Println("not a number")
    }
  }
}
