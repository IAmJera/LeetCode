package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node5 := ListNode{Val: 3, Next: nil}
	node4 := ListNode{Val: 3, Next: &node5}
	node3 := ListNode{Val: 2, Next: &node4}
	node2 := ListNode{Val: 1, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}
	head := deleteDuplicates(&node1)
	for ; head != nil; head = head.Next {
		fmt.Println(head.Val)
	}
}
func deleteDuplicates(head *ListNode) *ListNode {
	res := head
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return res
}
