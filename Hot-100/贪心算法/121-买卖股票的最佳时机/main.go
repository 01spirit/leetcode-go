package main

import "fmt"

func maxProfit(prices []int) int {
	minPrice := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		profit = max(profit, prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
	}
	return profit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
