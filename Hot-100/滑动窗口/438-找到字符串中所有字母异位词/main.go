package main

import (
	"fmt"
)

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}
	var sCount, pCount [26]int

	ans := make([]int, 0)
	for i := 0; i < len(p); i++ {
		pCount[p[i]-'a']++
		sCount[s[i]-'a']++
	}

	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i := len(p); i < len(s); i++ {
		sCount[s[i]-'a']++
		sCount[s[i-len(p)]-'a']--
		if sCount == pCount {
			ans = append(ans, i-len(p)+1)
		}
	}

	return ans
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}
