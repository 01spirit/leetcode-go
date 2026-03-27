package main

import "fmt"

func areSimilar(mat [][]int, k int) bool {
	n := len(mat)
	m := len(mat[0])
	rot := make([][]int, n)
	for i, arr := range mat {
		tmp := make([]int, m)
		for j, num := range arr {
			tmp[j] = num
		}
		rot[i] = tmp
	}

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			for j := 0; j < m; j++ {
				rot[i][j] = mat[i][(j+k*m+k)%m]
			}
		} else {
			for j := 0; j < m; j++ {
				rot[i][j] = mat[i][(j+k*m-k)%m]
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] != rot[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(areSimilar([][]int{{1, 2, 1, 2}, {5, 5, 5, 5}, {6, 3, 6, 3}}, 2))
	fmt.Println(areSimilar([][]int{{2, 2}, {2, 2}}, 3))
	fmt.Println(areSimilar([][]int{{1, 2}}, 1))
}
