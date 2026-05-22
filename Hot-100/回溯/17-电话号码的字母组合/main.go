package main

import "fmt"

var ref map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	n := len(digits)

	str := ""
	ans := []string{}

	var backtrack func(idx int, str string)
	backtrack = func(idx int, str string) {
		if idx == n {
			ans = append(ans, str)
			return
		}
		ch := digits[idx : idx+1]
		for _, c := range ref[ch] {
			str += string(c)
			backtrack(idx+1, str)
			str = str[:len(str)-1]
		}
	}

	backtrack(0, str)

	return ans
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations(""))
}
