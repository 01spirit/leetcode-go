package main

import "math"

type MinStack struct {
	stack  []int
	minEle []int
}

func Constructor() MinStack {
	return MinStack{
		stack:  []int{},
		minEle: []int{math.MaxInt32},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	top := this.minEle[len(this.minEle)-1]
	this.minEle = append(this.minEle, min(val, top))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minEle = this.minEle[:len(this.minEle)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minEle[len(this.minEle)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
