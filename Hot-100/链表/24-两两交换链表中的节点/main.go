package main

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

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	node0 := dummy
	node1 := head

	for node1 != nil && node1.Next != nil {
		node2 := node1.Next
		node0.Next = node2
		node1.Next = node2.Next
		node2.Next = node1

		node0 = node1
		node1 = node1.Next
	}

	return dummy.Next
}
