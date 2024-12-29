package main

import "fmt"

func canAliceWin(nums []int) bool {
	one := 0
	two := 0
	oth := 0

	for _, num := range nums {
		if num < 10 {
			one += num
		} else if num < 100 {
			two += num
		} else {
			oth += num
		}
	}

	return one > two+oth || two > one+oth
}

func main() {
	fmt.Println(canAliceWin([]int{1, 2, 3, 4, 10}))
	fmt.Println(canAliceWin([]int{1, 2, 3, 4, 5, 14}))
	fmt.Println(canAliceWin([]int{5, 5, 5, 25}))
}
