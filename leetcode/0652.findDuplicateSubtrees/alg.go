package _652_findDuplicateSubtrees

import (
	"strconv"
)

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	record := make(map[string]int)
	result := make([]*TreeNode, 0)
	findDuplicate(root, &result, record)
	return result
}

func findDuplicate(root *TreeNode, list *[]*TreeNode, record map[string]int) string {
	if root == nil {
		return "#"
	}
	left := findDuplicate(root.Left, list, record)
	right := findDuplicate(root.Right, list, record)
	r := strconv.FormatInt(int64(root.Val), 10)
	key := r + "," + left + "," + right
	record[key] += 1
	if record[key] == 2 {
		*list = append(*list, root)
	}
	return key
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
