package main

import (
	"math"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	res := 0

	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	str := ""
	length := 0
	for _, ch := range s {
		if !strings.Contains(str, string(ch)) {
			str += string(ch)
			length++
			//println(str)
		} else {
			idx := strings.Index(str, string(ch))

			res = int(math.Max(float64(res), float64(length)))

			str = str[idx+1:]
			length -= idx + 1
			str += string(ch)
			length++
			//println(str)
		}

	}
	res = int(math.Max(float64(res), float64(length)))
	return res
}

func main() {

	s := " "

	res := lengthOfLongestSubstring(s)

	println(res)
}
