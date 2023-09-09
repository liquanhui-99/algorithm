package algorithm

// TreeNode 定义树结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSymmetricTree 是否是对称二叉树
func isSymmetricTree(root *TreeNode) bool {
	return check(root, root)
}

func check(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}
	return check(left.Left, right.Right) && check(left.Right, right.Left)
}

// 题解
//对称二叉树的核心点：
//1. 树的右半部分是左半部分的镜像
//2. 同一层级的左树的左节点值应该等于右树的右节点值，左树的右节点应该等于右树的左节点
//3. 使用递归来实现层层遍历
//步骤：
//1. 判断同一层级的左树和右树同时为nil，索命是最后一层，不需要遍历了，且左右相等，停止递归
//2. 如果左树或者右树其中一个为nil，说明左右不想等，也不需要再比较下去了，停止递归
//3. 左右树都不为nil，那就对比左右树的节点值是否相等，如果不想等，那么左右树也不想等，停止递归
//4. 同一层级比较之后开始对比左树的左子节点是否和右树的右子节点相等，同时判断左树的右子节点是否和右树的左子节点
//相等，两个条件同时满足才是相等，需要如此递归下去。
