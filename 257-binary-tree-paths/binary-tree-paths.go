/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	var paths []string
	if root == nil {
		return paths
	}

	dfs(root, strconv.Itoa(root.Val), &paths)
	return paths
}

// dfs is a recursive helper function for depth-first search.
func dfs(node *TreeNode, path string, paths *[]string) {
	// Base case: If the node is a leaf, append the completed path to our results.
	if node.Left == nil && node.Right == nil {
		*paths = append(*paths, path)
		return
	}

	// Recursive step: Traverse left and right children if they exist.
	if node.Left != nil {
		dfs(node.Left, path+"->"+strconv.Itoa(node.Left.Val), paths)
	}
	if node.Right != nil {
		dfs(node.Right, path+"->"+strconv.Itoa(node.Right.Val), paths)
	}
}