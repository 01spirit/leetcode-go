package main

import "fmt"

func minDeletionSize(strs []string) int {
	n := len(strs[0])
	ans := 0

	for i := 0; i < n; i++ {
		for j := 0; j < len(strs)-1; j++ {
			if strs[j+1][i] < strs[j][i] {
				ans++
				break
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(minDeletionSize([]string{"cba", "daf", "ghi"}))
	fmt.Println(minDeletionSize([]string{"a", "b"}))
	fmt.Println(minDeletionSize([]string{"zyx", "wvu", "tsr"}))
}
