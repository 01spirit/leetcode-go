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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	left, right := dummy, dummy
	for ; n > 0; n-- {
		right = right.Next
	}
	for right.Next != nil {
		right = right.Next
		left = left.Next
	}
	left.Next = left.Next.Next

	return dummy.Next
}
