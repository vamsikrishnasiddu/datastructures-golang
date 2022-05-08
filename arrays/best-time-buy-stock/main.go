package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	buy := 0
	sell := 0

	profit := 0

	max := -9999

	for i := 0; i < len(prices); i++ {

		for j := i + 1; j < len(prices); j++ {

			if prices[j] > prices[i] {
				buy = prices[i]
				sell = prices[j]

				profit = sell - buy

				if profit > max {
					max = profit
				}

			}

		}
	}

	if max < 0 {
		return 0
	}

	return max

}

func maxProfitOptimal(prices []int) int {

	min := math.MaxInt64

	profit := 0

	for i := 0; i < len(prices); i++ {

		if prices[i] < min {
			min = prices[i]
		}

		if prices[i]-min > profit {
			profit = prices[i] - min
		}

	}

	return profit

}

func maxProfitOptimal2(prices []int) int {

	left := 0
	right := 1
	maxProfit := 0

	for right < len(prices) {
		if prices[left] > prices[right] {
			left = right
			right = right + 1
		} else if prices[left] <= prices[right] {

			if prices[right]-prices[left] > maxProfit {
				maxProfit = prices[right] - prices[left]

			}
			right++
		}
	}

	return maxProfit

}
func main() {

	var arr = []int{3, 3}

	profit := maxProfitOptimal2(arr)

	fmt.Println(profit)

}
