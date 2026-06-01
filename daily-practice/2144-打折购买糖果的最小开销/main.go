package main

import (
	"fmt"
	"sort"
)

func minimumCost(cost []int) int {
	sort.Ints(cost)
	n := len(cost)
	step := 3
	sum := 0
	for i := n - 1; i-step+1 >= 0; i -= step {
		sum += cost[i] + cost[i-1]
	}
	for i := n%step - 1; i >= 0; i-- {
		sum += cost[i]
	}
	return sum
}

func main() {
	fmt.Println(minimumCost([]int{1}))
	fmt.Println(minimumCost([]int{1, 2, 3}))
	fmt.Println(minimumCost([]int{6, 5, 7, 9, 2, 2}))
	fmt.Println(minimumCost([]int{5, 5}))
}
