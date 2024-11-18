package main

import "fmt"

func imageSmoother(img [][]int) [][]int {
	ans := make([][]int, len(img))

	for i := 0; i < len(img); i++ {
		rows := make([]int, len(img[i]))
		for j := 0; j < len(img[i]); j++ {
			cnt := 0
			sum := 0

			for row := i - 1; row <= i+1; row++ {
				if row < 0 || row >= len(img) {
					continue
				}
				for col := j - 1; col <= j+1; col++ {
					if col < 0 || col >= len(img[row]) {
						continue
					}
					cnt++
					sum += img[row][col]
				}
			}

			rows[j] = sum / cnt
		}
		ans[i] = rows
	}

	return ans
}

func main() {
	fmt.Println(imageSmoother([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
	fmt.Println(imageSmoother([][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}))
}
