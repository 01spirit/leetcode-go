package main

import (
	"fmt"
	"strconv"
)

func decodeString(s string) string {
	stack := []string{}
	ptr := 0
	for ptr < len(s) {
		ch := s[ptr]
		if ch >= '0' && ch <= '9' {
			digits := getDigit(s, &ptr)
			stack = append(stack, digits)
		} else if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '[' {
			stack = append(stack, string(ch))
			ptr++
		} else {
			ptr++
			str := []string{}
			for stack[len(stack)-1] != "[" {
				str = append(str, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			for i := 0; i < len(str)/2; i++ {
				str[i], str[len(str)-i-1] = str[len(str)-i-1], str[i]
			}
			stack = stack[:len(stack)-1]
			sub := getString(str)
			
			digits := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			num, _ := strconv.Atoi(digits)
			for i := 0; i < num; i++ {
				stack = append(stack, sub)
			}
		}
	}
	return getString(stack)
}

func getDigit(s string, ptr *int) string {
	digit := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		digit += string(s[*ptr])
	}
	return digit
}

func getString(arr []string) string {
	str := ""
	for _, s := range arr {
		str += s
	}
	return str
}

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("2[abc]3[cd]ef"))
	fmt.Println(decodeString("abc3[cd]xyz"))
}
