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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	jumpA, jumpB := false, false
	ptrA, ptrB := headA, headB
	for ptrA != nil && ptrB != nil {
		if ptrA == ptrB {
			return ptrA
		}
		if ptrA.Next == nil && !jumpA {
			ptrA = headB
			jumpA = true
		} else {
			ptrA = ptrA.Next
		}
		if ptrB.Next == nil && !jumpB {
			ptrB = headA
			jumpB = true
		} else {
			ptrB = ptrB.Next
		}
	}
	return nil
}
