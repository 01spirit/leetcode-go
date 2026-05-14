package main

import (
	"fmt"
	"slices"
)

func isGood(nums []int) bool {
	slices.Sort(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != i+1 {
			return false
		}
	}

	return len(nums)-1 == nums[len(nums)-1]
}

func main() {
	fmt.Println(isGood([]int{2, 1, 3}))
	fmt.Println(isGood([]int{1, 3, 3, 2}))
	fmt.Println(isGood([]int{1, 1}))
	fmt.Println(isGood([]int{3, 4, 4, 1, 2, 1}))
}
