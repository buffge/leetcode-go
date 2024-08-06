package main

import (
	"fmt"
	"github.com/buffge/leetcode-go/utils"
	"math"
)

type (
	TreeNode = utils.TreeNode
)

func dfs(node *TreeNode, left, right int) bool {
	if node == nil {
		return true
	}
	// 设置节点必须大于left 小于right (left,right)
	if node.Val <= left || node.Val >= right {
		return false
	}
	return dfs(node.Left, left, node.Val) && dfs(node.Right, node.Val, right)
}
func isValidBSTV1(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return dfs(root, math.MinInt64, math.MaxInt64)
}

/*
*
中序遍历
*/
func isValidBSTV2(root *TreeNode) bool {
	var stack []*TreeNode // 存放有左子节点的节点
	curr := root
	lastVal := math.MinInt64
	for len(stack) > 0 || curr != nil { // 中序遍历
		for curr != nil { // 找到最左
			stack = append(stack, curr)
			curr = curr.Left
		}
		node := stack[len(stack)-1] // 获取当前值
		stack = stack[:len(stack)-1]
		if node.Val <= lastVal {
			return false
		}
		lastVal = node.Val
		curr = node.Right
	}
	return true
}

var isValidBST = isValidBSTV1

func main() {
	root := &TreeNode{
		Val: 32,
		Left: &TreeNode{
			Val: 26,
			Left: &TreeNode{
				Val:   19,
				Right: &TreeNode{Val: 27},
			},
		},
		Right: &TreeNode{
			Val: 47,
			//Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 56},
		},
	}
	fmt.Println(isValidBST(root))
}
