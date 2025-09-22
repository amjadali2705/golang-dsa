/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    var result []int
    helperPreorder(root, &result)
    return result
}

func helperPreorder(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }

    *result = append(*result, node.Val)
    helperPreorder(node.Left, result)
    helperPreorder(node.Right, result)
}