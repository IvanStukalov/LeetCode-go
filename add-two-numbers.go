package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ptr1 := l1
	ptr2 := l2
	head := &ListNode{}
	ptr3 := head

	var carry int

	for ptr1 != nil && ptr2 != nil {
		node := &ListNode{
			Val:  carry,
			Next: nil,
		}

		if carry+ptr1.Val+ptr2.Val >= 10 {
			node.Val += ptr1.Val + ptr2.Val - 10
			carry = 1
		} else {
			node.Val += ptr1.Val + ptr2.Val
			carry = 0
		}

		ptr1 = ptr1.Next
		ptr2 = ptr2.Next

		ptr3.Next = node
		ptr3 = ptr3.Next
	}

	for ptr1 != nil {
		node := &ListNode{
			Val:  carry,
			Next: nil,
		}

		if carry+ptr1.Val >= 10 {
			node.Val += ptr1.Val - 10
			carry = 1
		} else {
			node.Val += ptr1.Val
			carry = 0
		}

		ptr1 = ptr1.Next

		ptr3.Next = node
		ptr3 = ptr3.Next
	}

	for ptr2 != nil {
		node := &ListNode{
			Val:  carry,
			Next: nil,
		}

		if carry+ptr2.Val >= 10 {
			node.Val += ptr2.Val - 10
			carry = 1
		} else {
			node.Val += ptr2.Val
			carry = 0
		}

		ptr2 = ptr2.Next

		ptr3.Next = node
		ptr3 = ptr3.Next
	}

	if carry != 0 {
		node := &ListNode{
			Val:  carry,
			Next: nil,
		}

		ptr3.Next = node
		ptr3 = ptr3.Next
	}

	return head.Next
}

func main() {
	node13 := &ListNode{
		Val:  9,
		Next: nil,
	}
	node12 := &ListNode{
		Val:  9,
		Next: node13,
	}
	node11 := &ListNode{
		Val:  9,
		Next: node12,
	}

	node23 := &ListNode{
		Val:  5,
		Next: nil,
	}

	list := addTwoNumbers(node11, node23)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
