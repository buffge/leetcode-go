package main

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newTree(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}
func check(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return check(left.Left, right.Right) && check(left.Right, right.Left)
}

// 是否为对称二叉树 递归版
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

// 是否为对称二叉树 loop
func isSymmetricV2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	q := []*TreeNode{root.Left, root.Right}
	for len(q) > 0 {
		left := q[0]
		right := q[1]
		q = q[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		q = append(q, left.Left, right.Right, left.Right, right.Left)
	}
	return true
}
func main() {
	t1 := newTree(1)
	t1.Left = newTree(2)
	t1.Left.Left = newTree(3)
	log.Println(isSymmetric(t1))
}
