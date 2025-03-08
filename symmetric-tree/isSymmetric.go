package symmetric_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return compareTrees(root.Left, root.Right)
}

func compareTrees(first *TreeNode, second *TreeNode) bool {
	if first == nil && second == nil {
		return true
	}
	if first == nil || second == nil {
		return false
	}
	if first.Val != second.Val {
		return false
	}

	return compareTrees(first.Left, second.Right) && compareTrees(first.Right, second.Left)
}
