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
递归法
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 如果root为空 忽略, 为p或q时 root为父节点
	if root == nil || root == p || root == q {
		return root
	}
	// 向下继续查找pq的祖先
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil { // 如果左边没找到 就在右边
		return right
	}
	if right == nil {
		return left
	}
	return root // 如果在左右都找到 父节点即为root
}

func main() {
	root := &TreeNode{Val: 10,
		Left: &TreeNode{Val: 5,
			Left: &TreeNode{Val: 3,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: -2},
			},
			Right: &TreeNode{Val: 2,
				Right: &TreeNode{Val: 1},
			},
		},
		Right: &TreeNode{Val: -3,
			Right: &TreeNode{Val: 11,
				Right: nil},
		},
	}
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Right))
}
