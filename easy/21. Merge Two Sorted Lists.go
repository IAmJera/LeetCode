package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node6 := ListNode{Val: 4, Next: nil}
	node5 := ListNode{Val: 3, Next: &node6}
	node4 := ListNode{Val: 1, Next: &node5}
	node3 := ListNode{Val: 4, Next: nil}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}

	head := mergeTwoLists(&node1, &node4)
	for ; head != nil; head = head.Next {
		fmt.Println(head.Val)
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	merged := &ListNode{}
	current := merged
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			current.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		current = current.Next
	}
	for list1 != nil {
		current.Next = list1
		break
	}
	for list2 != nil {
		current.Next = list2
		break
	}
	return merged.Next
}
