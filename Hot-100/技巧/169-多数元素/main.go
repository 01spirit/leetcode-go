package main

import "fmt"

func majorityElement(nums []int) int {
	hp := 0
	ans := 0
	for _, num := range nums {
		if hp == 0 {
			ans, hp = num, 1
		} else if ans == num {
			hp++
		} else {
			hp--
		}
	}
	return ans
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
