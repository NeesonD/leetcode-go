package _032_1_levelOrder

func levelOrder(root *TreeNode) []int {
	queue := make([]*TreeNode, 0)
	if root == nil {
		return []int{}
	}
	queue = append(queue, root)
	result := make([]int, 0)
	for len(queue) != 0 {
		le := len(queue)
		for i := 0; i < le; i++ {
			result = append(result, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[le:]
	}
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
