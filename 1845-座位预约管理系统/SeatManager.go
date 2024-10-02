package main

import (
	"container/heap"
	"fmt"
)

type SeatManager struct {
	available *Heap
}

func Constructor(n int) SeatManager {
	h := &Heap{}
	heap.Init(h)
	for i := 1; i <= n; i++ {
		heap.Push(h, i)
	}
	return SeatManager{available: h}
}

func (this *SeatManager) Reserve() int {
	return heap.Pop(this.available).(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(this.available, seatNumber)
}

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//type SeatManager struct {
//	seats []int
//	able  int
//}
//
//func Constructor(n int) SeatManager {
//	return SeatManager{seats: make([]int, n), able: 0}
//}
//
//func (this *SeatManager) Reserve() int {
//	ans := this.able + 1
//	this.seats[this.able] = 1
//	this.able++
//	for this.able < len(this.seats) && this.seats[this.able] != 0 {
//		this.able++
//	}
//
//	return ans
//}
//
//func (this *SeatManager) Unreserve(seatNumber int) {
//	this.seats[seatNumber-1] = 0
//	this.able = min(seatNumber-1, this.able)
//}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */

func main() {
	sm := Constructor(5)
	param1 := sm.Reserve()
	param2 := sm.Reserve()
	sm.Unreserve(2)
	param3 := sm.Reserve()
	param4 := sm.Reserve()
	param5 := sm.Reserve()
	param6 := sm.Reserve()
	sm.Unreserve(5)

	fmt.Println(param1, param2, param3, param4, param5, param6)
}
