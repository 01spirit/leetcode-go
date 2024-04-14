package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var tail *ListNode

	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0

		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + carry
		n := sum % 10
		carry = sum / 10

		if head == nil {
			head = &ListNode{
				n, nil,
			}
			tail = head
		} else {
			tail.Next = &ListNode{
				n, nil,
			}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return
}

func main() {
	l11 := ListNode{3, nil}
	l12 := ListNode{4, &l11}
	l13 := ListNode{2, &l12}

	l21 := ListNode{4, nil}
	l22 := ListNode{6, &l21}
	l23 := ListNode{5, &l22}

	res := addTwoNumbers(&l13, &l23)

	for res != nil {
		println(res.Val)
		res = res.Next
	}

}
