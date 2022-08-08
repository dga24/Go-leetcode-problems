package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
	len  int
}

func main() {
	linkedList := LinkedList{}
	linkedList.head = &ListNode{}
	linkedList.head.Val = 1
	linkedList.head.Next.Val = 2
	linkedList.head.Next.Next.Val = 3
	linkedList.head.Next.Next.Next.Val = 4
	linkedList.head.Next.Next.Next.Next.Val = 5

	fmt.Println(linkedList,"middleNode",middleNode(linkedList.head))

}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}