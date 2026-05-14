package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	sum := 0
	for cur := head.Next; cur != nil; cur = cur.Next {
		if cur.Val != 0 {
			sum += cur.Val
		} else {
			node := &ListNode{Val: sum}
			tail.Next = node
			tail = tail.Next
			sum = 0
		}
	}

	return dummy.Next
}

func main() {

}
