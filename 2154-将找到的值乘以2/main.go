package main

import (
	"fmt"
	"slices"
)

func findFinalValue(nums []int, original int) int {
	ans := original

	slices.Sort(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == ans {
			ans *= 2
		}
	}

	return ans
}

func main() {
	fmt.Println(findFinalValue([]int{5, 3, 6, 1, 12}, 3))
	fmt.Println(findFinalValue([]int{2, 7, 9}, 4))
}
