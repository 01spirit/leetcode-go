package main

import "fmt"

func isPossibleToSplit(nums []int) bool {
	seq := make([]int, 101)
	single := 0
	for i := 0; i < len(nums); i++ {
		if seq[nums[i]] == 2 {
			return false
		} else if seq[nums[i]] == 0 {
			single++
		} else if seq[nums[i]] == 1 {
			single--
		}
		seq[nums[i]]++
	}
	if single%2 != 0 {
		return false
	}

	return true
}

func main() {
	fmt.Println(isPossibleToSplit([]int{4, 4, 9, 10}))
	fmt.Println(isPossibleToSplit([]int{1, 1, 2, 2, 3, 4}))
	fmt.Println(isPossibleToSplit([]int{1, 1, 1, 1}))
}
