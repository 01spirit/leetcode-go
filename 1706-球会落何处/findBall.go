package main

import "fmt"

func findBall(grid [][]int) []int {
	n := len(grid[0])
	answer := make([]int, n)

	for i := range grid[0] {
		col := i
		for j, row := range grid {
			if row[col] == 1 && (col == n-1 || row[col+1] == -1) {
				answer[i] = -1
				break
			}
			if row[col] == -1 && (col == 0 || row[col-1] == 1) {
				answer[i] = -1
				break
			}
			if col == 0 && row[col] == -1 {
				answer[i] = -1
				break
			}
			col += row[col]
			if j == len(grid)-1 {
				answer[i] = col
			}
		}
	}

	return answer
}

func main() {
	fmt.Println(findBall([][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}}))
	fmt.Println(findBall([][]int{{-1}}))
	fmt.Println(findBall([][]int{{1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}, {1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}}))
}
