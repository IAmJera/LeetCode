package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node5 := ListNode{Val: 5, Next: nil}
	node4 := ListNode{Val: 4, Next: &node5}
	node3 := ListNode{Val: 3, Next: nil}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}
	//node5.Next = &node1
	node := addTwoNumbers(&node1, &node4)

	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
}

func getVal(node *ListNode) int {
	switch node {
	case nil:
		return 0
	default:
		return node.Val
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var out = new(ListNode)
	curr, accum, outVal := out, 0, 0
	for {
		cur := new(ListNode)
		outVal, cur.Next, accum = getVal(l1)+getVal(l2)+accum, nil, 0
		if outVal >= 10 {
			outVal = outVal - 10
			accum = 1
		}
		curr.Val, curr.Next = outVal, cur
		if l1.Next == nil && l2.Next == nil {
			switch accum {
			case 1:
				cur.Val, cur.Next = accum, nil
			default:
				curr.Next = nil
			}
			return out
		}

		switch {
		case l1.Next == nil:
			l1, l2, curr = new(ListNode), l2.Next, curr.Next
		case l2.Next == nil:
			l1, l2, curr = l1.Next, new(ListNode), curr.Next
		default:
			l2, l1, curr = l2.Next, l1.Next, curr.Next
		}
	}
}
