package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

var getIntersectionNode = getIntersectionNodeV1

/*
*
P99
双指针 让2个链表长度变成一样 从头开始判断是否一样 不一样都走一步
比如 A:7 :B:3
A走3步还剩4,B走3步还剩0 进入A
A走4步还剩0 进入B,B2走4步 还剩3 此时A2 B2都还剩3步 依次对比即可
*/
func getIntersectionNodeV1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB
	endTimes := 0
	for endTimes < 2 {
		pA = pA.Next
		pB = pB.Next
		if pA == nil {
			pA = headB
			endTimes++
		}
		if pB == nil {
			pB = headA
			endTimes++
		}
	}
	for pA.Next != nil && pB.Next != nil {
		if pA == pB {
			return pA
		}
		pA = pA.Next
		pB = pB.Next
	}
	if pA == pB {
		return pA
	}
	return nil
}

/*
*
优化版
*/
func getIntersectionNodeV2(headA, headB *ListNode) *ListNode {
	pA, pB := headA, headB
	// 2个都为空 表示 2个链表变成一样长 并且走到头都没有相同的子节点
	for pA != nil || pB != nil {
		// 用于将2个链表变成一样长
		if pA == nil {
			pA = headB
		}
		// 用于将2个链表变成一样长
		if pB == nil {
			pB = headA
		}
		// 如果节点相同则相交
		if pA == pB {
			return pA
		}
		pA = pA.Next
		pB = pB.Next
	}
	return nil
}
func main() {
	intersectNode := &ListNode{Val: 8}
	headA := &ListNode{Val: 1}
	headA.Next = &ListNode{Val: 2, Next: intersectNode}

	headB := &ListNode{Val: 10}
	headB.Next = &ListNode{
		Val: 11,
		Next: &ListNode{
			Val: 12,
			Next: &ListNode{
				Val: 13,
				Next: &ListNode{Val: 14,
					Next: intersectNode}}}}
	res := getIntersectionNodeV2(headA, headB)
	fmt.Println(res, res == intersectNode)
}
