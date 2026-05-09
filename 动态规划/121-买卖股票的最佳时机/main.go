package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	minPrice := prices[0]
	ans := 0
	for i := 0; i < len(prices); i++ {
		ans = max(ans, prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
	}
	return ans
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
