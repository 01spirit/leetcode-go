package main

import "fmt"

func numSpecial(mat [][]int) int {
	var ans int = 0

	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 1 {
				flag := true

				for ii := 0; ii < len(mat); ii++ {
					if mat[ii][j] == 1 && ii != i {
						flag = false
						break
					}
				}
				for jj := 0; jj < len(mat[i]); jj++ {
					if mat[i][jj] == 1 && jj != j {
						flag = false
						break
					}
				}

				if flag {
					ans++
				}
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(numSpecial([][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}))
	fmt.Println(numSpecial([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}))
}
