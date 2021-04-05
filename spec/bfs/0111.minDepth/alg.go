package _111_minDepth

import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	return int(math.Min(float64(minDepth(root.Left)), float64(minDepth(root.Right)))) + 1
}

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	depth := 1
	for len(q) != 0 {
		size := len(q)
		for _, node := range q {
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[size:]
		depth++
	}
	return depth
}
