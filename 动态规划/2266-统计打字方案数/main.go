package main

import "fmt"

const mod = 1e9 + 7

func substring(b byte, num int) int {
	dp := make([]int, num+1)
	dp[0] = 1
	dp[1] = 1
	if b-'7' == 0 || b-'9' == 0 {
		for i := 1; i <= num; i++ {
			if i >= 2 {
				dp[i] = (dp[i-1] + dp[i-2]) % mod
			}
			if i >= 3 {
				dp[i] = (dp[i] + dp[i-3]) % mod
			}
			if i >= 4 {
				dp[i] = (dp[i] + dp[i-4]) % mod
			}
		}
	} else {
		for i := 1; i <= num; i++ {
			if i >= 2 {
				dp[i] = (dp[i-1] + dp[i-2]) % mod
			}
			if i >= 3 {
				dp[i] = (dp[i] + dp[i-3]) % mod
			}
		}
	}
	return dp[num]
}

func countTexts(pressedKeys string) int {
	ans := 1

	start := 0
	for i := 1; i < len(pressedKeys); i++ {
		if pressedKeys[i-1] != pressedKeys[i] {
			ans = (ans * substring(pressedKeys[i-1], i-start)) % mod
			start = i
		}
	}
	ans = (ans * substring(pressedKeys[start], len(pressedKeys)-start)) % mod

	return ans % mod
}

func main() {
	fmt.Println(countTexts("22233"))
	fmt.Println(countTexts("222222222222222222222222222222222222"))
}
