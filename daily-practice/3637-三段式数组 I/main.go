package main

import "fmt"

const (
	zero  = 0
	one   = 1
	two   = 2
	three = 3
)

func isTrionic(nums []int) bool {
	status := 0

	for i := 0; i < len(nums)-1; i++ {
		switch status {
		case zero:
			if nums[i+1]-nums[i] > 0 {
				status = one
			} else {
				return false
			}
			break
		case one:
			if nums[i+1]-nums[i] > 0 {
				continue
			} else if nums[i+1]-nums[i] < 0 {
				status = two
			} else {
				return false
			}
			break
		case two:
			if nums[i+1]-nums[i] > 0 {
				status = three
			} else if nums[i+1]-nums[i] < 0 {
				continue
			} else {
				return false
			}
			break
		case three:
			if nums[i+1]-nums[i] > 0 {
				continue
			} else {
				return false
			}
		}
	}

	if status == three {
		return true
	}

	return false
}

func main() {
	fmt.Println(isTrionic([]int{1, 3, 5, 4, 2, 6}))
	fmt.Println(isTrionic([]int{2, 1, 3}))
}
