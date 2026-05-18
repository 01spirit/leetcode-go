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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	addition := 0
	ptr1, ptr2 := l1, l2
	head := &ListNode{}
	cur := head
	for ptr1 != nil && ptr2 != nil {
		val := ptr1.Val + ptr2.Val + addition
		if val > 9 {
			addition = 1
		} else {
			addition = 0
		}
		cur.Next = &ListNode{Val: val % 10}
		cur = cur.Next
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	for ptr1 != nil {
		val := ptr1.Val + addition
		if val > 9 {
			addition = 1
		} else {
			addition = 0
		}
		cur.Next = &ListNode{Val: val % 10}
		cur = cur.Next
		ptr1 = ptr1.Next
	}

	for ptr2 != nil {
		val := ptr2.Val + addition
		if val > 9 {
			addition = 1
		} else {
			addition = 0
		}
		cur.Next = &ListNode{Val: val % 10}
		cur = cur.Next
		ptr2 = ptr2.Next
	}

	if addition > 0 {
		cur.Next = &ListNode{Val: addition}
	}

	return head.Next
}
