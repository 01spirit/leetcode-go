package main

import "fmt"

func maxConsecutiveAnswers(answerKey string, k int) int {
	// 分别求 T 和 F 的最大长度
	return max(maxConsecutiveChar(answerKey, k, 'T'), maxConsecutiveChar(answerKey, k, 'F'))
}

func maxConsecutiveChar(answerKey string, k int, ch byte) int {
	left := 0
	count := 0
	num := 0

	for right := 0; right < len(answerKey); right++ {
		if answerKey[right] != ch {
			num++
		}
		for num > k {
			if answerKey[left] != ch {
				num--
			}
			left++
		}
		count = max(count, right-left+1)
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	answerKey := "TTFTTFTT"
	k := 1
	fmt.Println(maxConsecutiveAnswers(answerKey, k))
}
