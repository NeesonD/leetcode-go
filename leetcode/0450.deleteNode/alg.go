package _450_deleteNode

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		var getMinTreeNode func(node *TreeNode) *TreeNode
		getMinTreeNode = func(node *TreeNode) *TreeNode {
			for node.Left != nil {
				node = node.Left
			}
			return node
		}
		node := getMinTreeNode(root.Right)
		root.Val = node.Val
		root.Right = deleteNode(root.Right, node.Val)
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
