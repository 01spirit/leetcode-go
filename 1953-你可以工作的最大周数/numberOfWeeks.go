package main

import "fmt"

func numberOfWeeks(milestones []int) int64 {
	ans := int64(0)

	rest := int64(0)
	longest := 0
	for _, num := range milestones {
		longest = max(longest, num)
		rest += int64(num)
	}

	rest -= int64(longest)
	if rest < int64(longest-1) {
		ans = rest*2 + 1
	} else {
		ans = rest + int64(longest)
	}

	return ans
}

func main() {
	fmt.Println(numberOfWeeks([]int{1, 2, 3}))
	fmt.Println(numberOfWeeks([]int{1, 2, 5}))

}
