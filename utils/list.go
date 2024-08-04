package utils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Printf("->%d ", head.Val)
		head = head.Next
	}
	fmt.Println("")
}
