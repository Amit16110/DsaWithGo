package array

import (
	"fmt"
	"math"
)

func BestTimeBuyAndSell(price []int) int {
	lowestBuyPrice := price[0] + 1
	profitIfSoldToday := 0
	maxProfit := 0

	for _, priceToday := range price {
		if priceToday < lowestBuyPrice {
			lowestBuyPrice = priceToday
		}

		profitIfSoldToday = priceToday - lowestBuyPrice
		if profitIfSoldToday > maxProfit {
			maxProfit = profitIfSoldToday
		}
	}
	// Alternative solution
	// start loop with index 1.

	return maxProfit

}

// Buy and sell stock -2
//Maximum profit by buying and selling a share at most twice
//used two itration 1st => right to left and then 2nd => left to right.
// first used to get the maximum amount from the array.
// Second used to get the minmum amount from the array.

func MaxProfit(arr []int) {
	profit := []int{}
	n := len(arr)

	for i := 0; i < n; i++ {
		profit = append(profit, 0)

	}

	max_price := arr[n-1]

	for i := n - 2; i >= 0; i-- {
		if arr[i] > max_price {
			max_price = arr[i]
		}
		profit[i] = int(math.Max(float64(profit[i+1]), float64(max_price-arr[i])))
	}
	min_price := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < min_price {
			min_price = arr[i]
		}
		profit[i] = int(math.Max(float64(profit[i-1]), float64(profit[i]+(arr[i]-min_price))))
	}
	result := profit[n-1]
	fmt.Println("result", result)
}
