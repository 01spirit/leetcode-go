package main

import (
	"fmt"
	"slices"
)

func nextPermutation(nums []int) {
	n := len(nums)

	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	slices.Reverse(nums[i+1:])
}

func main() {
	fmt.Println(nextPermutation([]int{1, 2, 3}))
	fmt.Println(nextPermutation([]int{3, 2, 1}))
	fmt.Println(nextPermutation([]int{1, 1, 5}))
}
