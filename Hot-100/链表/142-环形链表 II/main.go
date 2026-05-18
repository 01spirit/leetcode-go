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

//func detectCycle(head *ListNode) *ListNode {
//	m := map[*ListNode]int{}
//	i := 0
//	node := head
//	for node != nil {
//		if _, ok := m[node]; ok {
//			return node
//		}
//		m[node] = i
//		i++
//		node = node.Next
//	}
//	return nil
//}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			cur := head
			for slow != cur {
				slow = slow.Next
				cur = cur.Next
			}
			return slow
		}
	}
	return nil
}
