package main

import (
	"fmt"
)

const mod = 1e9 + 7

func maxSubArray(nums []int) int {
	ans := 0
	preSum := 0
	minPreSsum := 0
	for _, num := range nums {
		preSum = preSum + num
		ans = max(ans, preSum-minPreSsum)
		minPreSsum = min(preSum, minPreSsum)
	}

	return ans
}

func kConcatenationMaxSum(arr []int, k int) int {
	if k == 1 {
		return maxSubArray(arr)
	}
	ans := maxSubArray(append(arr, arr...))
	sum := 0
	for _, num := range arr {
		sum += num
	}
	if sum > 0 {
		ans += sum * (k - 2)
	}
	return ans % mod
}

func main() {
	fmt.Println(maxSubArray([]int{-5, -2, 0, 0, 3, 9, -2, -5, 4}))
	fmt.Println(kConcatenationMaxSum([]int{1, 2}, 3))
	fmt.Println(kConcatenationMaxSum([]int{1, -2, 1}, 5))
	fmt.Println(kConcatenationMaxSum([]int{-1, -2}, 7))
}
