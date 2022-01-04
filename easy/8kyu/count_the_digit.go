/*
Take an integer n (n >= 0) and a digit d (0 <= d <= 9) as an integer.

Square all numbers k (0 <= k <= n) between 0 and n.

Count the numbers of digits d used in the writing of all the k**2.
*/

/*
  get int >= 0, int between 0 and 9
  return int

  for each number from 0 to first given, square the number
  from the squared numbers, count how many times the second given digit appears

  iterate from 0 to given, inclusive
    see how many times given digit is in number
    add to running total

  return total

  init total
  loop 0 to given incluseive
    get square of i
    count appearences
    add count to total
  return total

  countAppearences
    get int, int
    return how may times sencond int appears in first

    while number is above 0
      get % 10
      add 1 to count if equal to second arg
      reset number to number / 10


    init count to 0
    for loop
      init digit as number % 10
      if digit equals second arg
        add one to count
      set number to number % 10
      break if number is 0
    return count

*/

package main

import "fmt"

func countAppearances(number, digit int) int {
  if number == digit {
    return 1
  }

  count := 0

  for number != 0 {
    if number % 10 == digit {
      count++
    }
    number /= 10
  }

  return count
}

func NbDig(number, digit int) int {
  total := 0
  for i := 0; i <= number; i++ {
    square := i*i
    total += countAppearances(square, digit)
  }
  return total
}

func test(number, digit, expected int) {
  result := NbDig(number, digit)
  fmt.Println(result == expected)
}

func main() {
  test(10, 1, 4)
  test(25, 1, 11)
}
