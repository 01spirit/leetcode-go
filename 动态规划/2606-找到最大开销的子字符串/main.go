package main

import (
	"fmt"
	"math"
)

func maximumCostSubstring(s string, chars string, vals []int) int {
	prices := make([]int, 26)

	for i := range prices {
		prices[i] = math.MaxInt
	}
	for i := 0; i < len(chars); i++ {
		prices[chars[i]-'a'] = vals[i]
	}
	for i := 0; i < len(prices); i++ {
		if prices[i] == math.MaxInt {
			prices[i] = i + 1
		}
	}

	ans := 0
	preSum := 0
	minPreSum := 0
	for i := 0; i < len(s); i++ {
		preSum += prices[s[i]-'a']
		ans = max(ans, preSum-minPreSum)
		minPreSum = min(minPreSum, preSum)
	}

	return ans
}

func main() {
	fmt.Println(maximumCostSubstring("adaa", "d", []int{-1000}))
	fmt.Println(maximumCostSubstring("abc", "abc", []int{-1, -1, -1}))
}
