package main

import (
	"fmt"
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var copyRandomList = copyRandomListV1

func printNodeList(head *Node) {
	for head != nil {
		fmt.Printf("->(%d) %p [%p] ", head.Val, head, head.Random)
		if head.Random != nil {
			fmt.Printf("%d ", head.Random.Val)
		}
		head = head.Next
	}
	fmt.Println("")
}

/*
*
哈希表方法 存储random所在索引
*/
func copyRandomListV1(head *Node) *Node {
	nodeMapIdx := map[*Node]int{}
	nodes := make([]*Node, 0, 100)
	curr := head
	n := 0
	for curr != nil {
		nodes = append(nodes, curr)
		nodeMapIdx[curr] = n
		curr = curr.Next
		n++
	}
	nodes = append(nodes, nil)
	nodeMapIdx[nil] = n
	cloneNodes := make([]*Node, n+1)
	cloneNodes[n] = nil
	for i := n - 1; i >= 0; i-- {
		cloneNodes[i] = &Node{
			Val:  nodes[i].Val,
			Next: cloneNodes[i+1],
		}
	}
	for i := 0; i < n; i++ {
		cloneNodes[i].Random = cloneNodes[nodeMapIdx[nodes[i].Random]]
	}
	return cloneNodes[0]
}

/*
*
P99
优化版本 空间复杂度O(1)
复制加重新组合
第一轮循环 原链表扩大一倍  1 2 3 4 5 变成  11 22 33 44 55
第二轮循环 确定 copyNode的random节点 即 原node的next节点
第三轮循环 组合copyNode
*/
func copyRandomListV2(head *Node) *Node {
	if head == nil {
		return nil
	}
	curr := head
	for curr != nil { // 第一轮复制
		copyNode := &Node{
			Val:    curr.Val,
			Next:   curr.Next,
			Random: nil,
		}
		curr.Next = copyNode
		curr = copyNode.Next
	}
	curr = head
	for curr != nil { // 第二轮重建 random
		next := curr.Next
		if curr.Random != nil {
			next.Random = curr.Random.Next
		}
		curr = next.Next
	}
	res := head.Next
	curr = head
	copyCurr := head.Next
	for copyCurr.Next != nil { // 第三轮合并节点
		next := copyCurr.Next
		curr.Next = next          // 合并 原始链表
		copyCurr.Next = next.Next // 合并克隆链表
		curr = curr.Next
		copyCurr = copyCurr.Next
	}
	curr.Next = nil
	return res
}
func main() {
	head := &Node{
		Val:  7,
		Next: &Node{Val: 13, Next: &Node{Val: 11, Next: &Node{Val: 10, Next: &Node{Val: 1, Next: nil}}}},
	}
	head.Next.Random = head
	head.Next.Next.Random = head.Next.Next.Next.Next
	head.Next.Next.Next.Random = head.Next.Next
	head.Next.Next.Next.Next.Random = head
	//head = &Node{Val: 1}
	printNodeList(head)
	printNodeList(copyRandomListV2(head))
	printNodeList(head)

}
