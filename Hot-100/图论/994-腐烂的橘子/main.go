package main

import "fmt"

type Pair struct{ i, j int }

func orangesRotting(grid [][]int) int {
	cur := []Pair{}
	cnt := 0
	n, m := len(grid), len(grid[0])

	var check func(i, j int) bool
	check = func(i, j int) bool {
		if i < 0 || i > n-1 || j < 0 || j > m-1 || grid[i][j] == 0 || grid[i][j] == 2 {
			return false
		}
		grid[i][j] = 2
		return true
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 2 {
				cur = append(cur, Pair{i, j})
			}
		}
	}

	for len(cur) > 0 {
		tmp := []Pair{}
		for _, v := range cur {
			if check(v.i-1, v.j) {
				tmp = append(tmp, Pair{v.i - 1, v.j})
			}
			if check(v.i+1, v.j) {
				tmp = append(tmp, Pair{v.i + 1, v.j})
			}
			if check(v.i, v.j-1) {
				tmp = append(tmp, Pair{v.i, v.j - 1})
			}
			if check(v.i, v.j+1) {
				tmp = append(tmp, Pair{v.i, v.j + 1})
			}
		}

		cur = tmp
		if len(cur) > 0 {
			cnt++
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return cnt
}

func main() {
	fmt.Println(orangesRotting([][]int{{2, 2}, {1, 1}, {0, 0}, {2, 0}}))
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}))
	fmt.Println(orangesRotting([][]int{{0, 2}}))
}
