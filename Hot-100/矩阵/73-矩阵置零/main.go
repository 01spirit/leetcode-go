package main

func setZeroes(matrix [][]int) {
	row, col := map[int]bool{}, map[int]bool{}
	n, m := len(matrix), len(matrix[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for r, b := range row {
		if b {
			for j := 0; j < m; j++ {
				matrix[r][j] = 0
			}
		}
	}

	for c, b := range col {
		if b {
			for i := 0; i < n; i++ {
				matrix[i][c] = 0
			}
		}
	}

	//return matrix
}

func main() {
	//fmt.Println(setZeroes([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
	//fmt.Println(setZeroes([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}))
}
