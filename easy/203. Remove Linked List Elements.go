package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node7 := ListNode{Val: 6, Next: nil}
	node6 := ListNode{Val: 5, Next: &node7}
	node5 := ListNode{Val: 4, Next: &node6}
	node4 := ListNode{Val: 3, Next: &node5}
	node3 := ListNode{Val: 6, Next: &node4}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}
	head := removeElements(&node1, 6)
	for ; head != nil; head = head.Next {
		fmt.Println(head.Val)
	}

	node14 := ListNode{Val: 7, Next: nil}
	node13 := ListNode{Val: 7, Next: &node14}
	node12 := ListNode{Val: 7, Next: &node13}
	node11 := ListNode{Val: 7, Next: &node12}
	head1 := removeElements(&node11, 7)
	for ; head1 != nil; head1 = head1.Next {
		fmt.Println(head1.Val)
	}
}
func removeElements(head *ListNode, val int) *ListNode {
	prev := &ListNode{Next: head}
	res := prev
	for head != nil {
		if head.Val == val {
			prev.Next = head.Next
		} else {
			prev = head
		}
		head = head.Next
	}
	return res.Next
}
