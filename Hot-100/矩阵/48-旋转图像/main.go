package main

import "fmt"

func rotate(matrix [][]int) [][]int {
	n := len(matrix)
	for i := 0; i < (n+1)/2; i++ {
		for j := 0; j < n/2; j++ {
			matrix[n-1-j][i], matrix[n-1-i][n-1-j], matrix[j][n-1-i], matrix[i][j] = matrix[n-1-i][n-1-j], matrix[j][n-1-i], matrix[i][j], matrix[n-1-j][i]
		}
	}
	return matrix
}

func main() {
	fmt.Println(rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(rotate([][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}))
}
