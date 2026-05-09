package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	ans := math.MinInt
	preSum := 0
	minPreSum := 0
	for _, num := range nums {
		preSum += num
		ans = max(ans, preSum-minPreSum)
		minPreSum = min(preSum, minPreSum)
	}

	return ans
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
}
