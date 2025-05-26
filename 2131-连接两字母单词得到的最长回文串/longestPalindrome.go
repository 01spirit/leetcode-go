package main

import (
	"fmt"
	"strings"
)

func longestPalindrome(words []string) int {
	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}
	res := 0
	mid := false
	for word, cnt := range freq {
		rev := string(word[1]) + string(word[0])
		if word == rev {
			if cnt%2 == 1 {
				mid = true
			}
			res += 2 * (cnt / 2 * 2)
		} else if strings.Compare(word, rev) > 0 {
			res += 4 * min(freq[word], freq[rev])
		}
	}
	if mid {
		res += 2
	}
	return res
}

func main() {
	fmt.Println(longestPalindrome([]string{"lc", "cl", "gg"}))
	fmt.Println(longestPalindrome([]string{"ab", "ty", "yt", "lc", "cl", "ab"}))
	fmt.Println(longestPalindrome([]string{"cc", "ll", "xx"}))
}
