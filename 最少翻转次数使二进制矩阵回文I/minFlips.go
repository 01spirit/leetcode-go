package main

import "fmt"

func minFlips(grid [][]int) int {
	cnt := 0
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i])/2; j++ {
			if grid[i][j] != grid[i][len(grid[i])-1-j] {
				cnt++
			}
		}
	}
	ans = cnt
	cnt = 0
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid)/2; j++ {
			if grid[j][i] != grid[len(grid)-1-j][i] {
				cnt++
			}
		}
	}
	return min(ans, cnt)
}

func main() {
	fmt.Println(minFlips([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}))
	fmt.Println(minFlips([][]int{{0, 1}, {0, 1}, {0, 0}}))
	fmt.Println(minFlips([][]int{{1}, {0}}))
}
