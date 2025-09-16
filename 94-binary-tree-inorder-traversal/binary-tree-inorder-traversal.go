/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var result []int
    inorderHelper(root, &result)
    return result
}

func inorderHelper(node *TreeNode, result *[]int) {
    if node == nil {
        return 
    }

    inorderHelper(node.Left, result)
    *result = append(*result, node.Val)
    inorderHelper(node.Right, result)

}