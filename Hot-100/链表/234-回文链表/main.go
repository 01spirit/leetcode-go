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

func isPalindrome(head *ListNode) bool {
	arr := make([]int, 0)
	cur := head
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	l, r := 0, len(arr)-1
	for l < r {
		if arr[l] != arr[r] {
			return false
		}
		l++
		r--
	}
	return true
}
