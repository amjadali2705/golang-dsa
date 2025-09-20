/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }

    diff := maxDepth(root.Left) - maxDepth(root.Right)
    if diff <= 1 && diff >= -1 {
        return isBalanced(root.Left) && isBalanced(root.Right)
    }

    return false
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    leftDepth := maxDepth(root.Left)
    rightDepth := maxDepth(root.Right)

    return 1 + int(math.Max(float64(leftDepth), float64(rightDepth)))
}