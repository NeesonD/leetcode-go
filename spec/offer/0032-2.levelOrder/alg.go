package _032_2_levelOrder

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	if root == nil {
		return result
	}
	queue = append(queue, root)
	for len(queue) != 0 {
		le := len(queue)
		tmp := make([]int, 0)
		for i := 0; i < le; i++ {
			node := queue[i]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, tmp)
		queue = queue[le:]
	}
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
