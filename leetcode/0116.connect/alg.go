package _116_connect

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectHelp(root.Left, root.Right)
	return root
}

func connectHelp(node1 *Node, node2 *Node) {
	if node1 == nil {
		return
	}
	node1.Next = node2
	connectHelp(node1.Left, node1.Right)
	connectHelp(node1.Right, node2.Left)
	connectHelp(node2.Left, node2.Right)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
