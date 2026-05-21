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

/*
	自顶向下分治归并
*/
//func middleNode(head *ListNode) *ListNode {
//	pre := head
//	slow := pre
//	fast := pre
//	for fast != nil && fast.Next != nil {
//		pre = slow
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//	pre.Next = nil
//	return slow
//}
//
//func mergeTwoLists(l1, l2 *ListNode) *ListNode {
//	dummy := &ListNode{}
//	cur := dummy
//	for l1 != nil && l2 != nil {
//		if l1.Val < l2.Val {
//			cur.Next = l1
//			l1 = l1.Next
//		} else {
//			cur.Next = l2
//			l2 = l2.Next
//		}
//		cur = cur.Next
//	}
//	if l1 != nil {
//		cur.Next = l1
//	}
//	if l2 != nil {
//		cur.Next = l2
//	}
//	return dummy.Next
//}
//
//func sortList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//
//	mid := middleNode(head)
//	return mergeTwoLists(sortList(head), sortList(mid))
//}

/*
自底向上归并迭代
*/
func getListLength(head *ListNode) int {
	n := 0
	for head != nil {
		n++
		head = head.Next
	}
	return n
}

func split(head *ListNode, size int) *ListNode {
	cur := head
	for i := 0; i < size-1 && cur != nil; i++ {
		cur = cur.Next
	}
	if cur == nil || cur.Next == nil {
		return nil
	}
	nextHead := cur.Next
	cur.Next = nil

	return nextHead
}

func merge(l1, l2 *ListNode) (*ListNode, *ListNode) {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	for cur.Next != nil {
		cur = cur.Next
	}

	return dummy.Next, cur
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	length := getListLength(head)
	dummy := &ListNode{Next: head}
	for step := 1; step < length; step *= 2 {
		newListTail := dummy
		cur := dummy.Next
		for cur != nil {
			head1 := cur
			head2 := split(head1, step)
			cur = split(head2, step)
			head, tail := merge(head1, head2)
			newListTail.Next = head
			newListTail = tail
		}
	}
	return dummy.Next
}
