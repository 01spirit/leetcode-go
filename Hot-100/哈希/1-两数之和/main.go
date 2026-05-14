package main

import "fmt"

func twoSum(nums []int, target int) []int {
	n := len(nums)
	diff := make(map[int]int, n)
	for i := 0; i < n; i++ {
		if v, ok := diff[nums[i]]; ok {
			return []int{i, v}
		}
		diff[target-nums[i]] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
