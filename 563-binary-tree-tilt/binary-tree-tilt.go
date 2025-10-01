/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
    res := 0
    sum(root, &res)
    return res
}

func sum(node *TreeNode, tilt *int) int {
    if node == nil {
        return 0
    }

    leftSum := sum(node.Left, tilt)
    rightSum := sum(node.Right, tilt)

    diff := int(math.Abs(float64(leftSum - rightSum)))

    *tilt += diff

    return leftSum + rightSum + node.Val
}