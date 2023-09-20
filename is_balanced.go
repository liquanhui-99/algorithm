package algorithm

// IsBalanced 验证是否是平衡二叉树
func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return abs(height(root.Left)-height(root.Right)) <= 1 &&
		IsBalanced(root.Left) && IsBalanced(root.Right)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return max(height(node.Left), height(node.Right)) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
