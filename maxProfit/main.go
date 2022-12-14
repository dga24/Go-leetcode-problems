package main

import (
	"fmt"
	"math"
)

// You are given an array prices where prices[i] is the price of a given stock on the ith day.

// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0

func main() {
	prices := []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	min := math.MaxInt32
	max := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > max {
			max = prices[i] - min
		}
	}
	return max
}
