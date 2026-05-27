package main

import "fmt"

func exist(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])

	type Pair struct{ i, j int }
	var direction []Pair = []Pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	var check func(row, col, k int) bool
	check = func(row, col, k int) bool {
		if board[row][col] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		visited[row][col] = true
		defer func() { visited[row][col] = false }()
		for _, dir := range direction {
			if newRow, newCol := row+dir.i, col+dir.j; newRow >= 0 && newRow < n && newCol >= 0 && newCol < m && !visited[newRow][newCol] {
				if check(newRow, newCol, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"))
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"))
}
