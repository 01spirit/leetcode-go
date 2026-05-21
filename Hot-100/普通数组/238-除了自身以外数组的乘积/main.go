package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	suf := make([]int, n+1)
	suf[n] = 1
	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1] * nums[i]
	}

	pre := 1
	for i := 0; i < n; i++ {
		suf[i] = pre * suf[i+1]
		pre *= nums[i]
	}

	return suf[:n]
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}
