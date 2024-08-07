package main

import (
	"fmt"
	"github.com/buffge/leetcode-go/utils"
	"math"
)

type (
	TreeNode = utils.TreeNode
)

func dfs(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	left := max(0, dfs(root.Left, res))
	right := max(0, dfs(root.Right, res))
	*res = max(*res, root.Val+left+right)
	return root.Val + max(left, right)
}

/*
*
一个节点的最大路径和为 左中右3中相比 左右为负时设为0 全局最大值为 左中右,节点最大深度值为中+max(左,右)
*/
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := new(int)
	*res = math.MinInt
	dfs(root, res)
	return *res
}
func main() {
	root := &TreeNode{Val: 2,
		Left: &TreeNode{Val: -1}, //Left: &TreeNode{Val: 3,
		//	Left:  &TreeNode{Val: 3},
		//	Right: &TreeNode{Val: -2},
		//},
		//Right: &TreeNode{Val: 2,
		//	Right: &TreeNode{Val: 1},
		//},

		Right: &TreeNode{Val: -2}, //Right: &TreeNode{Val: 11, Right: nil},

	}
	fmt.Println(maxPathSum(root))
}
