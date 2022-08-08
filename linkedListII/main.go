package main

import "fmt"

func main ( ) {

    root := new ( ListNode )
    root.Insert ( 3 )
    root.Insert ( 2 )
	root.Insert (0)
	root.Insert(-4)

    root.Show()
}

type ListNode struct {
    Val int
    Next *ListNode
}

func (ln *ListNode) Insert(value int) {
    n := &ListNode {
        Val: value,
        Next: ln.Next,
    }
    ln.Next = n
}

func (ln *ListNode) Show() {
    for n := ln.Next; n != nil; n = n.Next {
        fmt.Println(n)
    }
}

func detectCycle(head *ListNode) *ListNode {
    
}

