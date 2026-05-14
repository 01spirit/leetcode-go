package main

import "fmt"

func findRotation(mat [][]int, target [][]int) bool {
	n := len(mat)

	if isEqual(mat, target) {
		return true
	}
	for k := 0; k < 4; k++ {
		for i := 0; i < n/2; i++ {
			for j := 0; j < (n+1)/2; j++ {
				mat[i][j], mat[n-1-j][i], mat[n-1-i][n-1-j], mat[j][n-1-i] = mat[n-1-j][i], mat[n-1-i][n-1-j], mat[j][n-1-i], mat[i][j]
			}
		}
		if isEqual(mat, target) {
			return true
		}
	}
	return false
}

func isEqual(mat [][]int, target [][]int) bool {
	n := len(mat)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != target[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(findRotation([][]int{{0, 1}, {1, 0}}, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(findRotation([][]int{[]int{0, 1}, []int{1, 1}}, [][]int{[]int{1, 0}, []int{0, 1}}))
	fmt.Println(findRotation([][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{1, 1, 1}}, [][]int{{1, 1, 1}, {0, 1, 0}, {0, 0, 0}}))
}
