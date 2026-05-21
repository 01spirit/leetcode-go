package main

import "fmt"

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	appear := make([]int, n+1)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		if i > 0 {
			ans[i] = ans[i-1]
		}
		appear[A[i]]++
		if appear[A[i]] == 2 {
			ans[i]++
		}
		appear[B[i]]++
		if appear[B[i]] == 2 {
			ans[i]++
		}
	}
	return ans
}

func main() {
	fmt.Println(findThePrefixCommonArray([]int{1, 3, 2, 4}, []int{3, 1, 2, 4}))
	fmt.Println(findThePrefixCommonArray([]int{2, 3, 1}, []int{3, 1, 2}))
}
