package main

func main() {

}

func reverseList(head *ListNode) *ListNode {
	var top ListNode
	if head == nil {
		return nil
	}
	top.Next = head

	for head.Next != nil {
		var next = head.Next
		head.Next = next.Next
		next.Next = top.Next
		top.Next = next
	}
	return top.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
