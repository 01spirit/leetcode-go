package main

import "fmt"

func check(nums []int) bool {
	n := len(nums)
	for i := 0; i < n; i++ {
		sorted := true
		for j := i; j < n+i-1; j++ {
			if nums[j%n] > nums[(j+1)%n] {
				sorted = false
				break
			}
		}
		if sorted {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(check([]int{3, 4, 5, 1, 2}))
	fmt.Println(check([]int{2, 1, 3, 4}))
	fmt.Println(check([]int{1, 2, 3}))
}
