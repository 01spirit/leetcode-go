package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type MinHeap []int

func (hp MinHeap) Len() int            { return len(hp) }
func (hp MinHeap) Less(i, j int) bool  { return hp[i] < hp[j] }
func (hp MinHeap) Swap(i, j int)       { hp[i], hp[j] = hp[j], hp[i] }
func (hp *MinHeap) Push(v interface{}) { *hp = append(*hp, v.(int)) }
func (hp *MinHeap) Pop() interface{}   { a := *hp; v := a[len(a)-1]; *hp = a[:len(a)-1]; return v }

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func solve(initScore, incScore []int) []int {
	mScore := make([]int, len(incScore)+1)

	hp := MinHeap{}
	for _, score := range initScore {
		hp = append(hp, score)
	}
	heap.Init(&hp)

	for _, score := range initScore {
		mScore[0] = max(mScore[0], score)
	}
	for i, inc := range incScore {
		minScore := heap.Pop(&hp).(int)
		curScore := minScore + inc
		mScore[i+1] = max(mScore[i], curScore)
		heap.Push(&hp, curScore)
	}
	return mScore[1:]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &m)

	initScore := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &initScore[i])
	}
	incScore := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &incScore[i])
	}

	ans := solve(initScore, incScore)

	for _, score := range ans {
		fmt.Fprintln(out, score)
	}
}
