package main

import "fmt"

func findPermutationDifference(s string, t string) int {
	pos := make(map[int32]int)
	diff := 0
	for i, ch := range s {
		pos[ch] = i
	}

	for i, ch := range t {
		cur := pos[ch]
		diff += abs(cur - i)
	}

	return diff
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	fmt.Println(findPermutationDifference("abc", "bac"))
	fmt.Println(findPermutationDifference("abcde", "edbca"))
}
