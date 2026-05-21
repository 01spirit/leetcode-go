package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	ans := math.MinInt32
	minPreSum := 0
	preSum := 0
	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		ans = max(ans, preSum-minPreSum)
		minPreSum = min(minPreSum, preSum)
	}
	return ans
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
}
