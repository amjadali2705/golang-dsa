/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    diameter := 0

    findDepth(root, &diameter)
    return diameter
}

func findDepth(node *TreeNode, diameter *int) int {
    if node == nil {
        return 0
    }

    left := findDepth(node.Left, diameter)
    right := findDepth(node.Right, diameter)

    if *diameter < left + right {
        *diameter = left + right
    }

    return 1 + int(math.Max(float64(left), float64(right)))
} 