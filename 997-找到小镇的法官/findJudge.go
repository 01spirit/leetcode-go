package main

import "fmt"

func findJudge(n int, trust [][]int) int {
	if n == 1 {
		if len(trust) == 0 {
			return 1
		}
		return -1
	}

	bt := make([]int, n+1)
	for i := 0; i < len(trust); i++ {
		bt[trust[i][0]] = -1
		if bt[trust[i][1]] != -1 {
			bt[trust[i][1]]++
		}
	}

	for i := range bt {
		if bt[i] == n-1 {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(findJudge(2, [][]int{{1, 2}}))
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}}))
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}))
	fmt.Println(findJudge(4, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}))
}
