package main

import (
	"fmt"
	"strings"
)

func isSubstringPresent(s string) bool {
	for i := len(s) - 1; i-1 >= 0; i-- {
		sub := []byte{s[i], s[i-1]}
		if strings.Contains(s, string(sub)) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isSubstringPresent("leetcode"))
	fmt.Println(isSubstringPresent("abcba"))
	fmt.Println(isSubstringPresent("abcd"))
}
