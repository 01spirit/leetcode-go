package main

import "fmt"

func permute(nums []int) [][]int {
	n := len(nums)
	res := [][]int{}
	path := []int{}
	used := make([]bool, n)

	var backtrack func()
	backtrack = func() {
		if len(path) == n {
			res = append(res, append([]int{}, path...))
			res = append(res, append([]int(nil), path...))
			return
		}

		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}

			used[i] = true
			path = append(path, nums[i])
			backtrack()
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	backtrack()

	return res
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{0, 1}))
	fmt.Println(permute([]int{1}))
}
