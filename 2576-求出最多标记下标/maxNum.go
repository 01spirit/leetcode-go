package main

import (
	"fmt"
	"slices"
)

func maxNumOfMarkedIndices(nums []int) int {
	count := 0

	if len(nums) == 1 {
		return count
	}

	slices.Sort(nums)
	small := len(nums)/2 - 1
	large := len(nums) - 1
	for small >= 0 {
		if nums[small]*2 <= nums[large] {
			count += 2
			small--
			large--
		} else {
			small--
		}
	}

	return count
}

func main() {
	fmt.Println(maxNumOfMarkedIndices([]int{3, 5, 2, 4}))
	fmt.Println(maxNumOfMarkedIndices([]int{9, 2, 5, 4}))
	fmt.Println(maxNumOfMarkedIndices([]int{7, 6, 8}))
}
