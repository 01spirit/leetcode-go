package main

import "fmt"

func countBinarySubstrings(s string) int {
	counts := []int{}
	ptr, n := 0, len(s)
	for ptr < n {
		c := s[ptr]
		count := 0
		for ptr < n && s[ptr] == c {
			ptr++
			count++
		}
		counts = append(counts, count)
	}
	ans := 0
	for i := 1; i < len(counts); i++ {
		ans += min(counts[i], counts[i-1])
	}
	return ans
}

func main() {
	fmt.Println(countBinarySubstrings("00110011"))
	fmt.Println(countBinarySubstrings("10101"))
}
