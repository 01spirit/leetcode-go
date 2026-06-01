package main

import "fmt"

func partition(s string) [][]string {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	ans := [][]string{}
	splits := []string{}
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]string{}, splits...))
			return
		}
		for j := idx; j < n; j++ {
			if f[idx][j] {
				splits = append(splits, s[idx:j+1])
				backtrack(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	backtrack(0)
	return ans
}

func main() {
	fmt.Println(partition("aab"))
	fmt.Println(partition("a"))
}
