package main

import (
	"fmt"
	"github.com/buffge/leetcode-go/utils"
)

type (
	TreeNode = utils.TreeNode
)

/*
中序遍历
*/
func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	curr := root
	for len(stack) > 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return node.Val
		}
		curr = node.Right
	}
	return 0
}
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
	fmt.Println(kthSmallest(root, 3))
}
