package main

import (
	"fmt"
	"math"
)

func closestTarget(words []string, target string, startIndex int) int {
	md := math.MaxInt32

	for i := 0; i < len(words); i++ {
		if words[i] == target {
			d := abs(i - startIndex)
			md = min(md, d, len(words)-d)
		}
	}
	if md == math.MaxInt32 {
		return -1
	}

	return md
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(closestTarget([]string{"hello", "i", "am", "leetcode", "hello"}, "hello", 1))
	fmt.Println(closestTarget([]string{"a", "b", "leetcode"}, "leetcode", 0))
	fmt.Println(closestTarget([]string{"i", "eat", "leetcode"}, "ate", 0))
}
