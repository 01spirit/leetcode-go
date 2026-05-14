package main

import "fmt"

func validStrings(n int) []string {
	str := ""
	res := make([]string, 0)
	backtrack(str, n, '0', &res)
	backtrack(str, n, '1', &res)

	return res
}

func backtrack(str string, n int, ch uint8, res *[]string) {
	if n == len(str) {
		if len(*res) == 0 || len(*res) > 0 && (*res)[len(*res)-1] != str {
			*res = append(*res, str)
		}
		return
	}

	if ch == '0' && len(str)-1 >= 0 && str[len(str)-1] == '0' {
		return
	}
	str += string(ch)
	backtrack(str, n, '0', res)
	backtrack(str, n, '1', res)
}

func main() {
	fmt.Println(validStrings(3))
	fmt.Println(validStrings(1))
}
