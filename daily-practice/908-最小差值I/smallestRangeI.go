package main

import (
	"fmt"
	"slices"
)

func smallestRangeI(nums []int, k int) int {
	//quickSort(nums)

	slices.Sort(nums)
	minEle, maxEle := nums[0], nums[len(nums)-1]
	if maxEle-minEle <= k*2 {
		return 0
	}

	return maxEle - minEle - k*2
}

func quickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	left, right := 0, len(nums)-1
	pivot := nums[right]
	lowerBound := 0
	for i := left; i < right; i++ {
		if nums[i] < pivot {
			tmp := nums[i]
			nums[i] = nums[lowerBound]
			nums[lowerBound] = tmp
			lowerBound++
		}
	}
	nums[right] = nums[lowerBound]
	nums[lowerBound] = pivot

	quickSort(nums[:lowerBound])
	quickSort(nums[lowerBound+1:])
}

func main() {
	fmt.Println(smallestRangeI([]int{1}, 0))
	fmt.Println(smallestRangeI([]int{0, 10}, 2))
	fmt.Println(smallestRangeI([]int{1, 3, 6}, 3))
}
