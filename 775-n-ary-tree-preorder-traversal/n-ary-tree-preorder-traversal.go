/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    var result []int
    preorderHelper(root, &result)
    return result
}

func preorderHelper(node *Node, result *[]int) {
    if node == nil {
        return
    }

    *result = append(*result, node.Val)
    
    for _, child := range node.Children {
        preorderHelper(child, result)
    }
}