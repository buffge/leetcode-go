package utils

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenerateTree(vals []*int) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	res := &TreeNode{Val: *vals[0]}
	stack := make([]*TreeNode, 0, len(vals))
	stack = append(stack, res)

	for len(stack) > 0 && len(vals) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(vals) > 0 {
			val := vals[0]
			vals = vals[1:]
			if val != nil {
				curr.Left = &TreeNode{Val: *val}
				stack = append(stack, curr.Left)
			}
		}
		if len(vals) > 0 {
			val := vals[0]
			vals = vals[1:]
			if val != nil {
				curr.Right = &TreeNode{Val: *val}
				stack = append(stack, curr.Right)
			}
		}
	}
	return res
}
