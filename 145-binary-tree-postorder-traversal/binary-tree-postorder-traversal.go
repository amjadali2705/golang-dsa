/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    var result []int
    postorderHelper(root, &result)
    return result
}

func postorderHelper(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }

    postorderHelper(node.Left, result)
    postorderHelper(node.Right, result)
    *result = append(*result, node.Val)

}