package main

import "fmt"

func countNegatives(grid [][]int) int {
	ans := 0
	m := len(grid)
	n := len(grid[0])

	end := n
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] < 0 {
				ans += (end - j) * (m - i)
				end = j
				break
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(countNegatives([][]int{[]int{4, 3, 2, -1}, []int{3, 2, 1, -1}, []int{1, 1, -1, -2}, []int{-1, -1, -2, -3}}))
	fmt.Println(countNegatives([][]int{[]int{3, 2}, []int{1, 0}}))
}
