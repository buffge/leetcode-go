package main

import (
	"fmt"
	"github.com/buffge/leetcode-go/utils"
)

type (
	TreeNode = utils.TreeNode
)

/*
*
dfs 遍历所有节点 遍历的时候更新最大深度
这里用了一个栈存深度 也可以用Val存 这样就会修改原数据
*/
func maxDepthV1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := []*TreeNode{root} // 存放待访问的节点
	depStack := []int{1}
	maxVal := 1
	for len(stack) > 0 {
		curr := stack[len(stack)-1]        // 获取当前节点
		stack = stack[:len(stack)-1]       // 此节点已被访问 删除
		depth := depStack[len(depStack)-1] // 获取当前节点的深度
		depStack = depStack[:len(depStack)-1]
		if curr.Left == nil && curr.Right == nil { // 如果当前节点没有子级 返回
			continue
		}
		maxVal = max(maxVal, depth+1) // 当有子节点是 子节点深度为depth+1 更新历史最大深度
		if curr.Left != nil {         // 将左子插入待访问栈
			stack = append(stack, curr.Left)
			depStack = append(depStack, depth+1)
		}
		if curr.Right != nil { // 将右子插入待访问栈
			stack = append(stack, curr.Right)
			depStack = append(depStack, depth+1)

		}
	}
	return maxVal
}

var maxDepth = maxDepthV1

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(maxDepth(root))
}
