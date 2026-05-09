package main

import (
	"fmt"
	"slices"
)

func maxProduct(nums []int) int {
	n := len(nums)
	fmax := make([]int, n)
	fmin := make([]int, n)
	fmax[0] = nums[0]
	fmin[0] = nums[0]
	for i := 1; i < n; i++ {
		fmax[i] = max(nums[i], fmax[i-1]*nums[i], fmin[i-1]*nums[i])
		fmin[i] = min(nums[i], fmax[i-1]*nums[i], fmin[i-1]*nums[i])
	}
	return slices.Max(fmax)
}

func main() {
	fmt.Println(maxProduct([]int{2, -5, -2, -4, 3}))
	fmt.Println(maxProduct([]int{-2, -3, 7}))
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
}
