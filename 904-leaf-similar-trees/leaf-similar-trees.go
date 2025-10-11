/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    list1 := make([]int, 0)
	list2 := make([]int, 0)

	leafSimilarHelper(root1, &list1)
	leafSimilarHelper(root2, &list2)

	return reflect.DeepEqual(list1, list2)
}

func leafSimilarHelper(node *TreeNode, leafList *[]int) {
    if node == nil {
        return
    }

    if node.Left == nil && node.Right == nil {
        *leafList = append(*leafList, node.Val)
        return
    }

    leafSimilarHelper(node.Left, leafList)
    leafSimilarHelper(node.Right, leafList)

}