package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head   *ListNode
	Length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head:   nil,
		Length: 0,
	}
}

func NewLinkedListFromSlice(nums []int) *LinkedList {
	list := NewLinkedList()
	for _, num := range nums {
		list.Append(num)
	}
	return list
}

func (l *LinkedList) Append(val int) {
	newNode := &ListNode{
		Val:  val,
		Next: nil,
	}
	if l.Head == nil {
		l.Head = newNode
		l.Length++
		return
	}
	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newNode
	l.Length++
}

func (l *LinkedList) Prepend(val int) {
	newNode := &ListNode{
		Val:  val,
		Next: nil,
	}
	if l.Head == nil {
		l.Head = newNode
		l.Length++
		return
	}
	next := l.Head.Next
	l.Head = newNode
	l.Head.Next = next
}

func (l *LinkedList) InsertAt(index, val int) {
	if index < 0 || index > l.Length {
		return
	}
	if index == 0 {
		l.Prepend(val)
		return
	}
	newNode := &ListNode{
		Val:  val,
		Next: nil,
	}
	cur := l.Head
	for i := 0; i < index-1; i++ {
		cur = cur.Next
	}
	newNode.Next = cur.Next
	cur.Next = newNode
	l.Length++
}

func (l *LinkedList) DeleteAt(index int) {
	if index < 0 || index > l.Length {
		return
	}
	if index == 0 {
		l.Head = l.Head.Next
		l.Length--
		return
	}
	cur := l.Head
	for i := 0; i < index-1; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	l.Length--
}

func (l *LinkedList) DeleteByValue(val int) {
	if l.Head == nil {
		return
	}
	if l.Head.Val == val {
		l.Head = l.Head.Next
		l.Length--
		return
	}
	cur := l.Head
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			l.Length--
			return
		}
		cur = cur.Next
	}
}

func (l *LinkedList) UpdateAt(index, val int) {
	if index < 0 || index > l.Length {
		return
	}
	cur := l.Head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.Val = val
}

func (l *LinkedList) UpdateByValue(oldVal, newVal int) {
	if l.Head == nil {
		return
	}
	cur := l.Head
	for cur != nil {
		if cur.Val == oldVal {
			cur.Val = newVal
			return
		}
		cur = cur.Next
	}
}
