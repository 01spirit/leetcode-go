package main

import "fmt"

func edgeScore(edges []int) int {
	score := make([]int64, len(edges))
	ans := -1
	ansScore := int64(0)

	for i := range edges {
		score[edges[i]] += int64(i)
		if score[edges[i]] > ansScore {
			ans = edges[i]
			ansScore = score[edges[i]]
		} else if score[edges[i]] == ansScore {
			ans = min(ans, edges[i])
		}
	}

	return ans
}

func main() {
	fmt.Println(edgeScore([]int{1, 0, 0, 0, 0, 7, 7, 5}))
	fmt.Println(edgeScore([]int{2, 0, 0, 2}))
	fmt.Println(edgeScore([]int{3, 3, 3, 0}))
}
