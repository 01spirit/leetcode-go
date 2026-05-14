package main

import "fmt"

func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {

	for i, i1 := x, x+k-1; i < i1; i, i1 = i+1, i1-1 {
		for j := y; j < y+k; j++ {
			grid[i][j], grid[i1][j] = grid[i1][j], grid[i][j]
		}
	}

	return grid
}

func main() {
	fmt.Println(reverseSubmatrix([][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}, []int{13, 14, 15, 16}}, 1, 0, 3))
	fmt.Println(reverseSubmatrix([][]int{{3, 4, 2, 3}, []int{2, 3, 4, 2}}, 0, 2, 2))
	fmt.Println(reverseSubmatrix([][]int{[]int{6, 16, 14}, []int{1, 2, 19}, []int{14, 17, 15}, []int{18, 7, 6}, []int{14, 12, 5}}, 2, 1, 2))
}
