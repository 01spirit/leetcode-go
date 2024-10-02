package main

import (
	"fmt"
)

func mincostTickets(days []int, costs []int) int {
	//lastDay := days[len(days)-1]
	//dp := make([]int, lastDay+1)
	//
	//minCost := slices.Min(costs)
	//minCostTH := min(costs[2], costs[0]*(30-lastDay-1)+costs[1]*4)
	//
	//index := 0
	//for i := 1; i <= lastDay && index < len(days); i++ {
	//	if days[index] != i {
	//		dp[i] = dp[i-1]
	//	} else {
	//		dp[i] = dp[i-1] + minCost
	//		if i-7 >= 0 {
	//			dp[i] = min(dp[i], dp[i-7]+costs[1])
	//		}
	//		if i-30 >= 0 {
	//			dp[i] = min(dp[i], dp[i-30]+costs[2])
	//		}
	//
	//		index++
	//	}
	//}
	//fmt.Println(dp)
	//
	//ans := min(minCostTH*int((lastDay+30)/30), dp[lastDay])
	//
	//return ans

	memo := [366]int{}
	dayM := map[int]bool{}
	for _, d := range days {
		dayM[d] = true
	}

	var dp func(day int) int
	dp = func(day int) int {
		if day > 365 {
			return 0
		}
		if memo[day] > 0 {
			return memo[day]
		}
		if dayM[day] {
			memo[day] = min(min(dp(day+1)+costs[0], dp(day+7)+costs[1]), dp(day+30)+costs[2])
		} else {
			memo[day] = dp(day + 1)
		}
		return memo[day]
	}
	return dp(1)
}

func main() {
	fmt.Println(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
	fmt.Println(mincostTickets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}))
	fmt.Println(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{7, 2, 15}))
	fmt.Println(mincostTickets([]int{1, 2, 3, 4, 6, 8, 9, 10, 13, 14, 16, 17, 19, 21, 24, 26, 27, 28, 29}, []int{3, 14, 50}))
	fmt.Println(mincostTickets([]int{1, 2, 4, 5, 6, 9, 10, 12, 14, 15, 18, 20, 21, 22, 23, 24, 25, 26, 28}, []int{3, 13, 57}))
}
