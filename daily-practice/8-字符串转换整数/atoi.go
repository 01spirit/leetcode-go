package main

import "math"

func myAtoi(s string) int {

	num := int64(0)
	flag := false
	neg := false
	for i, ch := range s {
		if ch == ' ' && !flag { // 前置空格
			continue
		} else {
			if flag && (ch < 48 || ch > 57) { // 数字后的无效字符
				break
			} else {
				if !flag && ch == '-' { // 负号
					neg = true
					flag = true
				} else if !flag && ch == '+' { // 正号，可能没有
					flag = true
					continue
				} else if ch >= 48 && ch <= 57 { // 数字
					flag = true
					num = num*10 + int64(ch-48)
					if num >= math.MaxInt32 && i < len(s)-1 && (s[i+1] >= 48 && s[i+1] <= 57) {
						break
					}
				} else { // 结束
					break
				}
			}
		}
	}

	if neg {
		num = -num
	}

	if num <= math.MinInt32 {
		num = math.MinInt32
	} else if num >= math.MaxInt32 {
		num = math.MaxInt32
	}

	return int(num)
}

func main() {
	s := "9223372036854775808"

	res := myAtoi(s)

	println(res)
}
