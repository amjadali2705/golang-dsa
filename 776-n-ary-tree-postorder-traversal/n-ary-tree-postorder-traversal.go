/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    var result []int
    postorderHelper(root, &result)
    return result
}

func postorderHelper(node *Node, result *[]int) {
    if node == nil {
        return
    }

    for _, child := range node.Children {
        postorderHelper(child, result)
    }

    *result = append(*result, node.Val)
}