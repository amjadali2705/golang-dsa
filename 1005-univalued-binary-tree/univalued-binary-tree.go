/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
    if root == nil {
        return true
    }

    return allNodesEqualVal(root.Left, root.Val) && allNodesEqualVal(root.Right, root.Val)
}

func allNodesEqualVal(node *TreeNode, val int) bool {
	if node == nil {
		return true
	}

	if node.Val != val {
		return false
	}

	return allNodesEqualVal(node.Left, val) && allNodesEqualVal(node.Right, val)
}