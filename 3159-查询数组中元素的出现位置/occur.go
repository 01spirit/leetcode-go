package main

import "fmt"

func occurrencesOfElement(nums []int, queries []int, x int) []int {
	targets := make([]int, len(nums))
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == x {
			targets[cnt] = i
			cnt++
		}
	}
	answer := make([]int, len(queries))
	idx := 0
	for _, num := range queries {
		if num > cnt {
			answer[idx] = -1
		} else {
			answer[idx] = targets[num-1]
		}
		idx++
	}

	return answer
}

func main() {
	fmt.Println(occurrencesOfElement([]int{1, 1, 3, 1, 1, 3, 2, 1}, []int{3}, 3))
	fmt.Println(occurrencesOfElement([]int{1, 3, 1, 7}, []int{1, 3, 2, 4}, 1))
	fmt.Println(occurrencesOfElement([]int{1, 2, 3}, []int{10}, 5))
}
