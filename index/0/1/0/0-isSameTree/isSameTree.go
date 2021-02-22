package main

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
思路：
递归
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
func newTree(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}
func main() {
	t1 := newTree(1)
	t1.Left = newTree(2)
	t1.Left.Left = newTree(3)
	t2 := newTree(1)
	t2.Left = newTree(2)
	log.Println(isSameTree(t1, t2))
}
