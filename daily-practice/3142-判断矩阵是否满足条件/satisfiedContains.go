package main

import "fmt"

func satisfiesConditions(grid [][]int) bool {
	for i := 0; i < len(grid)-1; i++ {
		for j := range grid[0] {
			if grid[i][j] != grid[i+1][j] {
				return false
			}
		}
	}

	for j := 0; j < len(grid[0])-1; j++ {
		if grid[0][j] == grid[0][j+1] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(satisfiesConditions([][]int{{1, 0, 2}, {1, 0, 2}}))
	fmt.Println(satisfiesConditions([][]int{{1, 1, 1}, {0, 0, 0}}))
	fmt.Println(satisfiesConditions([][]int{{1}, {2}, {3}}))
}
