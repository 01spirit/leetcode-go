package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {

	index := 0
	for index < len(gas) {
		if gas[index] < cost[index] {
			index++
			continue
		}
		tg := 0
		tc := 0
		for i := index; i < len(gas)+index; i++ {
			tg += gas[i%len(gas)]
			tc += cost[i%len(cost)]
			if tg < tc {
				index = i
				break
			} else {
				if i%len(gas) == (index+len(gas)-1)%len(gas) {
					return index
				}
			}
		}

	}

	return -1
}

func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
	fmt.Println(canCompleteCircuit([]int{3, 1, 1}, []int{1, 2, 2}))
}
