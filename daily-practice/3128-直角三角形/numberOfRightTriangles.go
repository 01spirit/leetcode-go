package main

func numberOfRightTriangles(grid [][]int) int64 {
	n := len(grid)
	nl := len(grid[0])
	rows := make([]int, n)
	cols := make([]int, nl)

	for i := 0; i < n; i++ {
		for j := 0; j < nl; j++ {
			if grid[i][j] == 1 {
				rows[i]++
				cols[j]++
			}
		}
	}

	ans := int64(0)

	for i := 0; i < n; i++ {
		for j := 0; j < nl; j++ {
			if grid[i][j] == 1 {
				ans += int64((rows[i] - 1) * (cols[j] - 1))
			}
		}

	}

	return ans
}
