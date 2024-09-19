package main

import "fmt"

func longestContinuousSubstring(s string) int {
	ans := 0

	for i := 0; i < len(s); i++ {
		cnt := 1
		for i < len(s)-1 && s[i+1]-s[i] == 1 {
			cnt++
			i++
		}
		ans = max(ans, cnt)
	}

	return ans
}

func main() {
	fmt.Println(longestContinuousSubstring("abacaba"))
	fmt.Println(longestContinuousSubstring("abcde"))
	fmt.Println(longestContinuousSubstring("a"))
}
