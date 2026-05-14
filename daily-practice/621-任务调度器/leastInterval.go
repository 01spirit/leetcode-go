package main

import "fmt"

func leastInterval(tasks []byte, n int) int {
	ans := 0

	cnt := make([]int, 26)
	for _, task := range tasks {
		cnt[task-'A']++
	}

	maxCnt := 0
	maxCntNum := 0
	for _, c := range cnt {
		if c == 0 {
			continue
		}
		if c > maxCnt {
			maxCnt = c
			maxCntNum = 1
		} else if c == maxCnt {
			maxCntNum++
		}
	}

	if (maxCnt-1)*(n+1)+maxCntNum > len(tasks) {
		ans = (maxCnt-1)*(n+1) + maxCntNum
	} else {
		ans = len(tasks)
	}

	return ans
}

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
	fmt.Println(leastInterval([]byte{'A', 'C', 'A', 'B', 'D', 'B'}, 1))
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 3))
}
