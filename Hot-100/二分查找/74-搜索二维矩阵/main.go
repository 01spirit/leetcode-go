package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])

	up, down := 0, n-1
	row := 0
	for up <= down {
		mid := up + (down-up)>>1
		if matrix[mid][0] <= target {
			row = mid
			up = mid + 1
		} else {
			down = mid - 1
		}
	}

	left, right := 0, m-1
	for left <= right {
		mid := left + (right-left)>>1
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1}, {3}, {5}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
}
