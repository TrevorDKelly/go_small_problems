/* Let us begin with an example:

A man has a rather old car being worth $2000. He saw a secondhand car being worth $8000. He wants to keep his old car until he can buy the secondhand one.

He thinks he can save $1000 each month but the prices of his old car and of the new one decrease of 1.5 percent per month. Furthermore this percent of loss increases of 0.5 percent at the end of every two months. Our man finds it difficult to make all these calculations.

Can you help him?

How many months will it take him to save up enough money to buy the car he wants, and how much money will he have left over? */

/*

  get int, int, int, float
  return int, int

  getting prices of two cars, additional savings per month, and can value % lostper month
  return months to save enough, $ left over

  BREAK DOWN
  track percent lost per month
    goes up .5 every two months
  track amount saved
    goes up 1K per month
  track car values
    go down by percent lost each month

  set
    available cash is 0 + oldcar value
    months = 0
    loss % is 1.5

  iterate till savings + old car value > new car value
    month + 1
    car values each down loss %
    break if $available > new car value
    if month is even
      add .5 to loss%

  ALGO
    cash available = oldval - float64
    months = 0
    loss percent = 1.5
    car vals to floats

    loop - index
      if cash available > new car val
        retrun months, int(cash available - new car val)
      months +1
      old car - old car val * percent lost
      new car - new car val * percent lost
      cash available + 1000
      if month % 2 == 0
        loss percent + 0.5


*/

package main

import (
  "fmt"
  "math"
)

func NbMonths(oldStartPrice, newStartPrice, savingsPerMonth int, lossPercent float64) [2]int {
	oldVal := float64(oldStartPrice)
	newVal := float64(newStartPrice)
	additionalCash := float64(savingsPerMonth)
	availableCash := 0.0
	months := 0

	for {
		if availableCash + oldVal >= newVal {
			return [2]int{months, int(math.Round(availableCash + oldVal - newVal))}
		}

		months++
		if months%2 == 0 {
			lossPercent += 0.5
		}
    oldVal -= (oldVal * (lossPercent / 100.0))
		newVal -= (newVal * (lossPercent / 100.0))
		availableCash += additionalCash
	}
}

func test(old, newcar, save int, loss float64, expected [2]int) {
	result := NbMonths(old, newcar, save, loss)
  fmt.Println(result, expected)
	fmt.Println(result[0] == expected[0] && result[1] == expected[1])
}

func main() {
	test(2000, 8000, 1000, 1.5, [2]int{6, 766})
	test(12000, 8000, 1000, 1.5, [2]int{0, 4000})
	test(8000, 12000, 500, 1.0, [2]int{8, 597})
	test(18000, 32000, 1500, 1.25, [2]int{8, 332})
	test(7500, 32000, 300, 1.55, [2]int{25, 122})
}
