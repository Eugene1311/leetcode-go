package construct_string_from_binary_tree

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(node *TreeNode) string {
	if node == nil {
		return ""
	}

	leftSubtreeStr := tree2str(node.Left)
	rightSubtreeStr := tree2str(node.Right)

	if len(leftSubtreeStr) > 0 {
		leftSubtreeStr = "(" + leftSubtreeStr + ")"
	}
	if len(leftSubtreeStr) == 0 && len(rightSubtreeStr) > 0 {
		leftSubtreeStr = "(" + leftSubtreeStr + ")"
	}

	if len(rightSubtreeStr) > 0 {
		rightSubtreeStr = "(" + rightSubtreeStr + ")"
	}
	return fmt.Sprintf("%d%s%s", node.Val, leftSubtreeStr, rightSubtreeStr)
}
