package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		node := &Node{Val: cur.Val, Next: cur.Next, Random: cur.Random}
		cur.Next = node
		cur = cur.Next.Next
	}
	ans := head.Next
	ptr := ans
	cur = head
	for ptr != nil && ptr.Next != nil {
		if ptr.Random != nil {
			ptr.Random = ptr.Random.Next
		}
		ptr = ptr.Next.Next
	}
	if ptr != nil && ptr.Random != nil {
		ptr.Random = ptr.Random.Next
	}
	ptr = ans
	for ptr != nil && ptr.Next != nil {
		next := ptr.Next
		ptr.Next = ptr.Next.Next
		ptr = ptr.Next
		cur.Next = next
		cur = cur.Next
	}
	cur.Next = nil
	return ans
}
