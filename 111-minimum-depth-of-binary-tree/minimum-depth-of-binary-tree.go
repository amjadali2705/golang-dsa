/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    if root.Left == nil && root.Right == nil {
        return 1
    }

    if root.Left == nil {
        return 1 + minDepth(root.Right)
    }

    if root.Right == nil {
        return 1 + minDepth(root.Left)
    }

    leftDepth := minDepth(root.Left)
    rightDepth := minDepth(root.Right)

    return 1 + int(math.Min(float64(leftDepth), float64(rightDepth)))
}

