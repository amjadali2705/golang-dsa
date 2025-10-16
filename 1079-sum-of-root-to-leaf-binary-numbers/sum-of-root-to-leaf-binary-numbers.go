/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
    return dfs(root, 0)
}

func dfs(node *TreeNode, currentNum int) int {
    if node == nil {
        return 0
    }

    currentNum = (currentNum << 1) | node.Val

    if node.Left == nil && node.Right == nil {
        return currentNum
    }

    return dfs(node.Left, currentNum) + dfs(node.Right, currentNum)
}