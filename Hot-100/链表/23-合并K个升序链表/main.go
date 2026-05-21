package main

import "container/heap"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type MinHeap []*ListNode

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *MinHeap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func mergeKLists(lists []*ListNode) *ListNode {
	hp := MinHeap{}
	dummy := &ListNode{}
	cur := dummy
	for _, list := range lists {
		if list != nil {
			hp = append(hp, list)
		}
	}
	heap.Init(&hp)
	for len(hp) > 0 {
		v := heap.Pop(&hp).(*ListNode)
		if v.Next != nil {
			heap.Push(&hp, v.Next)
		}
		cur.Next = v
		cur = cur.Next
	}
	return dummy.Next
}
