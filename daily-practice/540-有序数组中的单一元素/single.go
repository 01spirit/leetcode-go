package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if mid%2 == 0 {
			if mid-1 >= 0 && nums[mid] == nums[mid-1] {
				right = mid - 1
			} else if mid+1 < len(nums) && nums[mid] == nums[mid+1] {
				left = mid + 1
			} else {
				return nums[mid]
			}
		} else {
			if mid-1 >= 0 && nums[mid] == nums[mid-1] {
				left = mid + 1
			} else if mid+1 < len(nums) && nums[mid] == nums[mid+1] {
				right = mid - 1
			} else {
				return nums[mid]
			}
		}

	}

	return nums[left]
}

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3}))
	fmt.Println(singleNonDuplicate([]int{1, 1, 2}))
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	fmt.Println(singleNonDuplicate([]int{1, 1, 3, 3, 4, 4, 5, 8, 8}))
	fmt.Println(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))
	fmt.Println(singleNonDuplicate([]int{3, 3, 4, 7, 7, 11, 11}))
}
