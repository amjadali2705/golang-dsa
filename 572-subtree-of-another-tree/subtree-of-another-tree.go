/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil || subRoot == nil {
		return root == subRoot
	}

	if root.Val == subRoot.Val && isIdentical(root, subRoot) {
		return true
	}
    
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isIdentical(s, t *TreeNode) bool {
	if s == nil || t == nil {
		return s == t
	}

	return s.Val == t.Val && isIdentical(s.Left, t.Left) && isIdentical(s.Right, t.Right)
}