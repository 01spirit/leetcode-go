package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	hash := map[byte]int{}
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i > 0 {
			delete(hash, s[i-1])
		}
		for rk+1 < n && hash[s[rk+1]] == 0 {
			hash[s[rk+1]]++
			rk++
		}
		ans = max(ans, rk+1-i)
	}
	return ans
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abba"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("a"))
	fmt.Println(lengthOfLongestSubstring("aab"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
