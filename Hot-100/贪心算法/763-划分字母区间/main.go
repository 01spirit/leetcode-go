package main

import "fmt"

func partitionLabels(s string) []int {
	last := map[byte]int{}
	for i := 0; i < len(s); i++ {
		last[s[i]] = i
	}

	ans := []int{}
	for i := 0; i < len(s); i++ {
		end := last[s[i]]
		for j := i + 1; j < end; j++ {
			if last[s[j]] > end {
				end = last[s[j]]
			}
		}
		ans = append(ans, end-i+1)
		i = end
	}
	return ans
}

func main() {
	fmt.Println(partitionLabels("eaaaabaaec"))
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels("eccbbbbdec"))
}
