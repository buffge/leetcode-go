package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("->%d ", head.Val)
		head = head.Next
	}
	fmt.Println("")
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	curr := res
	overflow := 0
	for {
		if l1 != nil {
			curr.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			curr.Val += l2.Val
			l2 = l2.Next
		}
		curr.Val += overflow
		if curr.Val > 9 {
			overflow = 1
			curr.Val -= 10
		} else {
			overflow = 0
		}
		if l1 == nil && l2 == nil {
			if overflow == 1 {
				curr.Next = &ListNode{Val: 1}
			}
			break
		}
		curr.Next = &ListNode{}
		curr = curr.Next
	}
	return res
}
func main() {
	l1 := &ListNode{
		Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 7, Next: nil}}},
	}
	l2 := &ListNode{
		Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: nil}}},
	}
	printList(addTwoNumbers(l1, l2))
}
