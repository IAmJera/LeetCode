package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node5 := ListNode{Val: 5, Next: nil}
	node4 := ListNode{Val: 4, Next: &node5}
	node3 := ListNode{Val: 3, Next: &node4}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}

	head := reverseList(&node1)
	for ; head != nil; head = head.Next {
		fmt.Println(head.Val)
	}
}

func reverseList(head *ListNode) *ListNode {
	var past *ListNode = nil
	for head != nil {
		head.Next, past, head = past, head, head.Next
	}
	return past
}
