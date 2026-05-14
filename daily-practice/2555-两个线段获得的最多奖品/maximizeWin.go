package main

import (
	"fmt"
	"sort"
)

func maximizeWin(prizePositions []int, k int) int {
	n := len(prizePositions)
	dp := make([]int, n+1)
	ans := 0
	for i := 0; i < n; i++ {
		j := sort.SearchInts(prizePositions, prizePositions[i]-k)
		ans = max(ans, i-j+1+dp[j])
		dp[i+1] = max(dp[i], i-j+1)
	}
	return ans
}

func main() {
	fmt.Println(maximizeWin([]int{1, 1, 2, 2, 3, 3, 5}, 2))
	fmt.Println(maximizeWin([]int{1, 2, 3, 4}, 0))
}
