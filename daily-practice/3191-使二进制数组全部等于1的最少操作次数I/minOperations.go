package main

import "fmt"

func minOperations(nums []int) int {
	ans := 0
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == 0 {
			ans++
			nums[i] = 1
			nums[i+1] ^= 1
			nums[i+2] ^= 1
		}
	}
	n := len(nums)
	if nums[n-1] == 0 || nums[n-2] == 0 {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(minOperations([]int{0, 1, 1, 1, 0, 0}))
	fmt.Println(minOperations([]int{0, 1, 1, 1}))
}
