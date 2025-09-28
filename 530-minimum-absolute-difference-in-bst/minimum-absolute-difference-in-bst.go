/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
    minDiff := math.MaxInt64
    prevVal := -1

    inOrderTraversal(root, &minDiff, &prevVal)

    return minDiff
}


func inOrderTraversal(node *TreeNode, minDiff *int, prevVal *int) {
    if node == nil {
        return
    }

    inOrderTraversal(node.Left, minDiff, prevVal)

    if *prevVal != -1 {
        diff := int(math.Abs(float64(node.Val - *prevVal)))
        if diff < *minDiff {
            *minDiff = diff
        }
    }

    *prevVal = node.Val

    inOrderTraversal(node.Right, minDiff, prevVal)
}