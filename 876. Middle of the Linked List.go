package main

import "fmt"

func main() {
	node6 := ListNode{Val: 5, Next: nil}
	node5 := ListNode{Val: 5, Next: &node6}
	node4 := ListNode{Val: 4, Next: &node5}
	node3 := ListNode{Val: 3, Next: &node4}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}
	fmt.Println(middleNode(&node1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

//func middleNode(head *ListNode) *ListNode {
//	var nodes []*ListNode
//	for head != nil {
//		nodes = append(nodes, head)
//		head = head.Next
//	}
//	return nodes[len(nodes)/2]
//}
