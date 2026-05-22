package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	upBound, lowBound, leftBound, rightBound := 0, n-1, 0, m-1

	for i := 0; i < n; i++ {
		if matrix[i][m-1] == target {
			return true
		} else if matrix[i][m-1] > target {
			upBound = i
			break
		}
	}
	for i := n - 1; i >= 0; i-- {
		if matrix[lowBound][0] == target {
			return true
		} else if matrix[lowBound][0] < target {
			lowBound = i
			break
		}
	}
	for i := 0; i < m; i++ {
		if matrix[lowBound][i] == target {
			return true
		} else if matrix[lowBound][i] < target {
			leftBound = i
			break
		}
	}
	for i := m - 1; i >= 0; i-- {
		if matrix[upBound][i] == target {
			return true
		} else if matrix[upBound][i] < target {
			rightBound = i
			break
		}
	}

	for i := upBound; i <= lowBound; i++ {
		left, right := leftBound, rightBound
		for left <= right {
			mid := (left + right) / 2
			if matrix[i][mid] == target {
				return true
			} else if matrix[i][mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))
	fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20))
}
