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
		minPreSum = min(minPreSum, preSum)
	}
	return ans
}

func minSubArray(nums []int) int {
	ans := math.MaxInt
	preSum := 0
	maxPreSum := 0
	for _, num := range nums {
		preSum += num
		ans = min(ans, preSum-maxPreSum)
		maxPreSum = max(maxPreSum, preSum)
	}
	return ans
}

func maxSubarraySumCircular(nums []int) int {
	maxSum := maxSubArray(nums)
	minSum := minSubArray(nums)

	if maxSum < 0 {
		return maxSum
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	return max(maxSum, sum-minSum)
}

func main() {
	fmt.Println(maxSubarraySumCircular([]int{1, -2, 3, -2}))
	fmt.Println(maxSubarraySumCircular([]int{5, -3, 5}))
	fmt.Println(maxSubarraySumCircular([]int{3, -2, 2, -3}))
}
