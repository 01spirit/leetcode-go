package main

import (
	"container/heap"
	"fmt"
)

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	if len(stations) == 0 {
		if target > startFuel {
			return -1
		}
		return 0
	}
	h := new(maxHeap)
	heap.Init(h)

	ans := 0
	costFuel := 0
	haveFuel := startFuel
	for i := 0; i < len(stations); i++ {
		if target <= stations[i][0] {
			return ans
		}
		costFuel = stations[i][0]

		for h.Len() > 0 && costFuel > haveFuel {
			haveFuel += heap.Pop(h).(int)
			ans++
		}
		if h.Len() == 0 && costFuel > haveFuel {
			return -1
		}
		heap.Push(h, stations[i][1])
	}

	for h.Len() > 0 && target > haveFuel {
		haveFuel += heap.Pop(h).(int)
		ans++
	}

	if target > haveFuel {
		return -1
	}

	return ans
}

func main() {
	fmt.Println(minRefuelStops(100, 50, [][]int{{25, 30}}))
	fmt.Println(minRefuelStops(200, 50, [][]int{{25, 25}, {50, 100}, {100, 100}, {150, 40}}))
	fmt.Println(minRefuelStops(100, 1, [][]int{}))
	fmt.Println(minRefuelStops(1, 1, [][]int{}))
	fmt.Println(minRefuelStops(100, 1, [][]int{{10, 100}}))
	fmt.Println(minRefuelStops(100, 10, [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}))
}
