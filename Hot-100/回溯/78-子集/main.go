package main

import "fmt"

func subsets(nums []int) [][]int {
	n := len(nums)
	res := [][]int{}
	path := []int{}

	var backtrack func(int)
	backtrack = func(idx int) {
		res = append(res, append([]int{}, path...))
		if idx == n {
			return
		}
		for i := idx; i < n; i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)

	return res
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{0}))
}
