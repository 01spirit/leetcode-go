package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	res := []string{}

	var backtrack func(path string, left, right int)
	backtrack = func(path string, left, right int) {
		if right == n {
			res = append(res, path)
			return
		}
		if left < n {
			backtrack(path+"(", left+1, right)
		}
		if right < left {
			backtrack(path+")", left, right+1)
		}
	}

	backtrack("", 0, 0)

	return res
}

func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(1))
}
