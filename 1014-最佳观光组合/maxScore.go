package main

import "fmt"

func maxScoreSightseeingPair(values []int) int {
	ans := 0
	mscore := values[0] + 0
	for i := 1; i < len(values); i++ {
		ans = max(ans, mscore+values[i]-i)
		mscore = max(mscore, values[i]+i)
	}

	return ans
}

func main() {
	fmt.Println(maxScoreSightseeingPair([]int{8, 1, 5, 2, 6}))
	fmt.Println(maxScoreSightseeingPair([]int{1, 2}))
	fmt.Println(maxScoreSightseeingPair([]int{1, 2, 2}))
}
