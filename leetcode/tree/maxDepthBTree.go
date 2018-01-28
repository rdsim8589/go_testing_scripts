package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftVal := maxDepth(root.Left)
	rightVal := maxDepth(root.Right)

	if rightVal > leftVal {
		return rightVal + 1
	}
	return leftVal + 1
}
