package main

import "fmt"

func countPartitions(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	left := 0
	right := sum
	cnt := 0
	for i := 0; i < len(nums)-1; i++ {
		left += nums[i]
		right -= nums[i]
		if (left-right)%2 == 0 {
			cnt++
		}
	}

	return cnt
}

func main() {
	fmt.Println(countPartitions([]int{10, 10, 3, 7, 6}))
	fmt.Println(countPartitions([]int{1, 2, 2}))
	fmt.Println(countPartitions([]int{2, 4, 6, 8}))
}
