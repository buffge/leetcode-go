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

/*
*
双指针 记录上一个节点 同时更新上一个节点和当前节点
*/
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func main() {
	head := &ListNode{
		Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
	}
	printList(head)
	printList(reverseList(head))
}
