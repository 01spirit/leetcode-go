package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	res := [][]int{}
	path := []int{}

	var backtrack func(sum int, idx int)
	backtrack = func(sum int, idx int) {
		if sum == target {
			res = append(res, append([]int{}, path...))
			return
		}
		if sum > target {
			return
		}
		for i := idx; i < n; i++ {
			path = append(path, candidates[i])
			backtrack(sum+candidates[i], i)
			path = path[:len(path)-1]
		}
	}

	backtrack(0, 0)

	return res
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{2}, 1))
}
