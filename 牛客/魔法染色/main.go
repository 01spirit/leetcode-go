package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	var s string
	fmt.Fscan(in, &n, &m, &s)

	// 二分答案 k（区间长度上限）
	// 标准答案下界取 0，因为可能全是 R，不需要施法
	left, right := 0, n
	ans := n
	for left <= right {
		mid := (left + right) >> 1
		if check(s, n, m, mid) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	fmt.Fprintln(out, ans)
}

// check：当每次施法长度上限为 k 时，能否用不超过 m 次施法染红全墙
func check(s string, n, m, k int) bool {
	// dp[i] 表示前 i 个位置（1-indexed 意义下）染红的最小次数
	dp := make([]int, n+1)
	// dp[0] = 0 已满足

	for i := 1; i <= n; i++ {
		if s[i-1] == 'R' { // Go string 是 0-indexed
			// 已经是红色，继承前 i-1 个位置的结果
			dp[i] = dp[i-1]
		} else {
			// s[i-1] == 'W'
			// 必须在 i 结束一次施法，区间长度不超过 k
			prev := i - k
			if prev < 0 {
				prev = 0
			}
			dp[i] = dp[prev] + 1
		}
		// 如果已经超过 m，提前剪枝（可选）
		if dp[i] > m {
			return false
		}
	}

	return dp[n] <= m
}
