package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(arr []int) int {
	ans := 0
	n := len(arr)
	leftEnd, rightStart := 0, n-1
	for leftEnd < n-1 && arr[leftEnd] <= arr[leftEnd+1] {
		leftEnd++
	}
	if leftEnd == n-1 {
		return n * (n + 1) / 2
	}
	for rightStart > 0 && arr[rightStart-1] <= arr[rightStart] {
		rightStart--
	}

	ans += n - rightStart + 1

	j := rightStart
	for i := 0; i <= leftEnd; i++ {
		if j < i+2 {
			j = i + 2
		}
		for j < n && arr[i] > arr[j] {
			j++
		}
		ans += n - j + 1
	}

	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	ans := solve(arr)

	fmt.Fprintln(out, ans)
}
