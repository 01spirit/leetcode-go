package main

import "fmt"

func plusOne(digits []int) []int {
	flag := false

	ans := make([]int, len(digits))
	ans[len(ans)-1] = (digits[len(ans)-1] + 1) % 10
	if digits[len(ans)-1] == 9 {
		flag = true
	}
	for i := len(digits) - 2; i >= 0; i-- {
		if flag {
			if digits[i] == 9 {
				flag = true
			} else {
				flag = false
			}
			ans[i] = (digits[i] + 1) % 10
		} else {
			flag = false
			ans[i] = digits[i]
		}
	}

	if flag {
		ans = append([]int{1}, ans...)
	}

	return ans
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{8, 9, 9, 9}))
}
