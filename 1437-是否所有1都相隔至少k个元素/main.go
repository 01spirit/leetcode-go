package main

import "fmt"

func kLengthApart(nums []int, k int) bool {
	start := 0
	begin := false
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			if begin && i-start <= k {
				return false
			}
			if !begin {
				begin = true
			}
			start = i
		}
	}

	return true
}

func main() {
	fmt.Println(kLengthApart([]int{1, 0, 0, 0, 1, 0, 0, 1}, 2))
	fmt.Println(kLengthApart([]int{1, 0, 0, 1, 0, 1}, 2))
}
