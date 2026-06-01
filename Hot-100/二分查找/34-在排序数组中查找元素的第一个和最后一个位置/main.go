package main

import "fmt"

func searchRange(nums []int, target int) []int {
	n := len(nums)
	left, right := 0, n-1
	st, ed := -1, -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			st, ed = mid, mid
			break
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if st == -1 {
		return []int{-1, -1}
	}
	for st >= 0 && nums[st] == target {
		st--
	}
	for ed <= n-1 && nums[ed] == target {
		ed++
	}
	return []int{st + 1, ed - 1}
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{}, 0))
}
