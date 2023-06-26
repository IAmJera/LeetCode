package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//node6 := ListNode{Val: 5, Next: nil}
	node5 := ListNode{Val: 5, Next: nil}
	node4 := ListNode{Val: 4, Next: &node5}
	node3 := ListNode{Val: 3, Next: &node4}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}
	//node5.Next = &node1
	fmt.Println(hasCycle(&node1))
}
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
