package main

import "fmt"

func resultsArray(nums []int, k int) []int {
	results := make([]int, len(nums)-k+1)
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i]-nums[i-1] != 1 {
			cnt = 1
		} else {
			cnt++
		}
		if cnt >= k {
			results[i-k+1] = nums[i]
		} else if i-k+1 >= 0 {
			results[i-k+1] = -1
		}
	}
	return results
}

func main() {
	fmt.Println(resultsArray([]int{1, 2, 3, 4, 3, 2, 5}, 3))
	fmt.Println(resultsArray([]int{2, 2, 2, 2, 2}, 4))
	fmt.Println(resultsArray([]int{3, 2, 3, 2, 3, 2}, 2))
}
