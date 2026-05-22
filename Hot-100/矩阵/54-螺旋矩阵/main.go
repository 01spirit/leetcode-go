package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	ans := []int{}
	n, m := len(matrix), len(matrix[0])

	rowLim, colLim := 0, 0
	for len(ans) < n*m {
		for j := colLim; j < m-colLim; j++ {
			ans = append(ans, matrix[rowLim][j])
		}
		for i := rowLim + 1; i < n-rowLim; i++ {
			ans = append(ans, matrix[i][m-colLim-1])
		}
		for j := m - colLim - 2; j >= colLim; j-- {
			ans = append(ans, matrix[n-rowLim-1][j])
		}
		for i := n - rowLim - 2; i > rowLim; i-- {
			ans = append(ans, matrix[i][colLim])
		}
		rowLim++
		colLim++
	}

	return ans[:n*m]
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}
