package main

import "fmt"

func longestConsecutive(nums []int) int {
	has := map[int]bool{}
	for _, n := range nums {
		has[n] = true
	}

	ans := 0
	for x := range has {
		if has[x-1] {
			continue
		}
		y := x + 1
		for has[y] {
			y++
		}
		ans = max(ans, y-x)
		if ans*2 >= len(nums) {
			break
		}
	}
	return ans
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive([]int{1, 0, 1, 2}))
}
